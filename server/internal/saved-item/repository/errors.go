package repository

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrDatabase = errors.New("database error")
)

func MapError(err error) error {
	return err
}
