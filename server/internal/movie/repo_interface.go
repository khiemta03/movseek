package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetOneMovie(ctx context.Context, movieID int64) (models.Movie, error)
	ListMovies(ctx context.Context, input ListMoviesOptions) ([]models.Movie, error)
	CountMovies(ctx context.Context, input ListMoviesOptions) (int, error)
	GetUpcomingMovies(ctx context.Context, input GetUpcomingMoviesOptions) ([]models.MovieSummary, error)
	CountUpcomingMovies(ctx context.Context, input GetUpcomingMoviesOptions) (int, error)
	GetTrendingMovies(ctx context.Context, input GetTrendingMoviesOptions) ([]models.MovieSummary, error)
	CountTrendingMovies(ctx context.Context, input GetTrendingMoviesOptions) (int, error)
	GetTopRatedMovies(ctx context.Context, input GetTopRatedMoviesOptions) ([]models.MovieSummary, error)
	CountTopRatedMovies(ctx context.Context, input GetTopRatedMoviesOptions) (int, error)
	GetNowPlayingMovies(ctx context.Context, input GetNowPlayingMoviesOptions) ([]models.MovieSummary, error)
	CountNowPlayingMovies(ctx context.Context, input GetNowPlayingMoviesOptions) (int, error)
	GetPopularMovies(ctx context.Context, input GetPopularMoviesOptions) ([]models.MovieSummary, error)
	CountPopularMovies(ctx context.Context, input GetPopularMoviesOptions) (int, error)
	GetMovieGenres(ctx context.Context) ([]models.MovieGenre, error)
	GetLastestTrailer(ctx context.Context, input GetLastestTrailerInput) ([]GetLastestTrailerResponse, error)
}
