package repository

import (
	"errors"

	"github.com/tmplam/movseek/pkg/mongo"
)

var (
	ErrNotFound = errors.New("not found")
	ErrDatabase = errors.New("database error")
)

func MapError(err error) error {
	switch err {
	case mongo.ErrNoDocuments:
		return ErrNotFound
	}
	return ErrDatabase
}
