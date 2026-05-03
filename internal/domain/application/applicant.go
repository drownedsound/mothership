package application

import (
	"errors"
	"time"
	"unicode/utf8"
)

type Applicant struct {
	Last  string
	First string
	DOB   time.Time
}

func NewApplicant(l string, f string, d time.Time) (*Applicant, error) {
	a := &Applicant{
		Last:  l,
		First: f,
		DOB:   d,
	}

	if err := a.Validate(); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Applicant) Validate() error {
	var errs []error

	if utf8.RuneCountInString(a.Last) <= 0 {
		errs = append(errs, errors.New("missing last name"))
	}

	if utf8.RuneCountInString(a.First) <= 0 {
		errs = append(errs, errors.New("missing first name"))
	}

	if a.DOB.IsZero() {
		errs = append(errs, errors.New("missing date of birth"))
	}

	return errors.Join(errs...)
}
