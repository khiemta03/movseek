package mongo

import (
	"time"

	"github.com/tmplam/movseek/internal/person"
	"github.com/tmplam/movseek/pkg/mongo"
)

type implRepository struct {
	db    mongo.Database
	clock func() time.Time
}

func NewRepository(db mongo.Database) person.Repository {
	return &implRepository{
		db:    db,
		clock: time.Now,
	}
}
