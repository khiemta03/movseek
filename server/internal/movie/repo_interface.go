package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetOneMovie(ctx context.Context, movieID string) (models.Movie, error)
	ListMovies(ctx context.Context, input ListMoviesOptions) ([]models.Movie, error)
	GetUpcomingMovies(ctx context.Context, input GetUpcomingMoviesOptions) ([]models.MovieSummary, error)
	GetTrendingMovies(ctx context.Context, input GetTrendingMoviesOptions) ([]models.MovieSummary, error)
	GetTopRatedMovies(ctx context.Context, input GetTopRatedMoviesOptions) ([]models.MovieSummary, error)
	GetPopularMovies(ctx context.Context, input GetPopularMoviesOptions) ([]models.MovieSummary, error)
}
