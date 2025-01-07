package mongo

import (
	"time"

	"github.com/tmplam/movseek/internal/tvshow"
	"github.com/tmplam/movseek/pkg/mongo"
)

type implRepository struct {
	db    mongo.Database
	clock func() time.Time
}

func NewRepository(db mongo.Database) tvshow.Repository {
	return &implRepository{
		db:    db,
		clock: time.Now,
	}
}
