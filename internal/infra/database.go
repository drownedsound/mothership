package infra

import (
	"errors"
	"os"
)

var ErrMissingDSN = errors.New("missing dsn")

type Database struct {
	DSN string
}

func NewDatabase(d string) (*Database, error) {
	if d == "" {
		d = os.Getenv("MOT_DSN")
	}

	if d == "" {
		return nil, ErrMissingDSN
	}

	return &Database{DSN: d}, nil
}
