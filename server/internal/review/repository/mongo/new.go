package mongo

import (
	"time"

	"github.com/tmplam/movseek/internal/review"
	"github.com/tmplam/movseek/pkg/mongo"
)

type implRepository struct {
	db    mongo.Database
	clock func() time.Time
}

func NewRepository(db mongo.Database) review.Repository {
	return &implRepository{
		db:    db,
		clock: time.Now,
	}
}
