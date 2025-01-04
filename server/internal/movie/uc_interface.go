package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Usecase interface {
	GetOneMovie(ctx context.Context, movieID string) (models.Movie, error)
	GetMovieCredits(ctx context.Context, movieID string) (models.MovieCredits, error)
	ListMovies(ctx context.Context, input ListMoviesInput) (ListMoviesOutput, error)
}
