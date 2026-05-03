package main

import (
	"errors"
	"fmt"
	"time"

	app "github.com/drownedsound/mothership/internal/domain/application"
	"github.com/drownedsound/mothership/internal/infra"
)

const (
	lineLength = 80
	filler     = "-"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from unexpected panic")
		}
	}()
	printHeader()
	defer printFooter()

	johnSmith, _ := app.NewApplicant("Smith", "John", time.Date(1979, 8, 30, 0, 0, 0, 0, time.UTC))
	creditCard, err := app.NewApplication(johnSmith, app.ProductCreditCard, 1_000_000)

	if err == nil {
		fmt.Printf("Product == %v\n", creditCard.Product)
		fmt.Printf("Limit == %d\n", creditCard.Limit)
		fmt.Printf("Last Name == %s\n", creditCard.Last)
		fmt.Printf("First Name == %s\n", creditCard.First)
		fmt.Printf("Date of Birth == %s\n", creditCard.DOB.Format(time.DateOnly))
	} else {
		fmt.Println(err)
	}

	db, err := infra.NewDatabase("this/is/some/path/")
	if errors.Is(err, infra.ErrMissingDSN) {
		panic(fmt.Errorf("unable to initialize database: %w", err))
	}
	fmt.Printf("DSN == %s\n", db.DSN)

	s := infra.NewServer(infra.LoggerAdapter(logToConsole))
	fmt.Printf("server == %+v\n\n", s)
}
