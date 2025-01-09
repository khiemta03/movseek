package mongo

import (
	"time"

	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"github.com/tmplam/movseek/pkg/mongo"
)

type implRepository struct {
	db    mongo.Database
	clock func() time.Time
}

func NewRepository(db mongo.Database) saved_item.Repository {
	return &implRepository{
		db:    db,
		clock: time.Now,
	}
}
