package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Usecase interface {
	GetOneMovie(ctx context.Context, movieID int64) (models.Movie, error)
	GetMovieCredits(ctx context.Context, movieID int64) (models.MovieCredits, error)
	GetMovieVideos(ctx context.Context, movieID int64) ([]models.Trailer, error)
	GetMovieKeywords(ctx context.Context, movieID int64) ([]models.KeyWord, error)
	ListMovies(ctx context.Context, input ListMoviesInput) (ListMoviesOutput, error)
	GetUpcomingMovies(ctx context.Context, input GetUpcomingMoviesInput) (GetUpcomingMoviesOutput, error)
	GetTrendingMovies(ctx context.Context, input GetTrendingMoviesInput) (GetTrendingMoviesOutput, error)
	GetTopRatedMovies(ctx context.Context, input GetTopRatedMoviesInput) (GetTopRatedMoviesOutput, error)
	GetPopularMovies(ctx context.Context, input GetPopularMoviesInput) (GetPopularMoviesOutput, error)
	GetNowPlayingMovies(ctx context.Context, input GetNowPlayingMoviesInput) (GetNowPlayingMoviesOutput, error)
	GetMovieGenres(ctx context.Context) (GetMovieGenresOutput, error)
	GetLastestTrailer(ctx context.Context, input GetLastestTrailerInput) (GetLastestTrailerOutput, error)
}
