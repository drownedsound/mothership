package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
	"errors"
)

// TODO: Use enums and const
// TODO: Embed structs
// TODO: Wrap errors
// TODO: Use errors.Is and errors.As
// TODO: Use sentinel errors for missing config file
// TODO: Use implicit interfaces to inject dependencies
// TODO: Use defer, panic and recover
// TODO: Use stringbuilder

type applicant struct {
	last string
	first string
	dob time.Time
}

func newApplicant(l string, f string, d time.Time) (*applicant, error) {
	a := &applicant {
		last: l,
		first: f,
		dob: d,
	}
	
	if err := a.validate(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *applicant) validate() error {
	var errs []error

	if utf8.RuneCountInString(a.last) <= 0 {
		errs = append(errs, errors.New("missing last name"))
	}

	if utf8.RuneCountInString(a.first) <= 0 {
		errs = append(errs, errors.New("missing first name"))
	}

	if a.dob.IsZero() {
		errs = append(errs, errors.New("missing date of birth"))
	}

	return errors.Join(errs...)
}

const lineLength = 80
const filler = "-"

func printHeader () {
	fmt.Println(strings.Repeat(filler, lineLength))	
	fmt.Println(strings.ToUpper("👾 👾 👾 Start Mothership 👾 👾 👾"))
	fmt.Println(strings.Repeat(filler, lineLength))	
}

func printFooter () {
	fmt.Println(strings.Repeat(filler, lineLength))	
	fmt.Println(strings.ToUpper("👾 👾 👾 End Mothership 👾 👾 👾"))
	fmt.Println(strings.Repeat(filler, lineLength))	
}

func main() {
	printHeader()
	defer printFooter()

	johnSmith, err := newApplicant("Smith", "John", time.Date(1979, 8, 30, 0, 0, 0, 0, time.UTC))
	// johnSmith, err := newApplicant("", "", time.Time{})
	if err == nil {
		fmt.Printf("Last Name == %s\n", johnSmith.last)
		fmt.Printf("First Name == %s\n", johnSmith.first)
		fmt.Printf("Date of Birth== %s\n", johnSmith.dob.Format(time.DateOnly))
	} else {
		fmt.Println(err)
	}
}
