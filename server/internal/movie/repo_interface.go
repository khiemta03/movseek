package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetOneMovie(ctx context.Context, movieID string) (models.Movie, error)
}
