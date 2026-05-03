package application

import "errors"

type Product uint8

const (
	ProductCreditCard = iota
	ProductPersonalLoan
	ProductMax
)

type Application struct {
	*Applicant
	Product Product
	Limit   int
}

func NewApplication(a *Applicant, p Product, l int) (*Application, error) {
	app := &Application{
		Applicant: a,
		Product:   p,
		Limit:     l,
	}

	if err := app.Validate(); err != nil {
		return nil, err
	}
	return app, nil
}

func (a *Application) Validate() error {
	var errs []error

	if a.Product >= ProductMax {
		errs = append(errs, errors.New("invalid product"))
	}

	return errors.Join(errs...)
}
