package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie"
)

func (uc implUsecase) GetOneMovie(ctx context.Context, movieID string) (models.Movie, error) {
	return uc.repo.GetOneMovie(ctx, movieID)
}

func (uc implUsecase) GetMovieCredits(ctx context.Context, movieID string) (models.MovieCredits, error) {
	m, err := uc.repo.GetOneMovie(ctx, movieID)
	if err != nil {
		return models.MovieCredits{}, err
	}

	return m.Credits, nil
}

func (uc implUsecase) ListMovies(ctx context.Context, input movie.ListMoviesInput) (movie.ListMoviesOutput, error) {
	movies, err := uc.repo.ListMovies(ctx, movie.ListMoviesOptions{
		Query:  input.Query,
		Filter: input.Filter,
	})
	if err != nil {
		return movie.ListMoviesOutput{}, err
	}

	return movie.ListMoviesOutput{Movies: movies}, nil
}

func (uc implUsecase) GetUpcomingMovies(ctx context.Context, input movie.GetUpcomingMoviesInput) (movie.GetUpcomingMoviesOutput, error) {
	movies, err := uc.repo.GetUpcomingMovies(ctx, movie.GetUpcomingMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetUpcomingMoviesOutput{}, err
	}

	return movie.GetUpcomingMoviesOutput{Movies: movies}, nil
}

func (uc implUsecase) GetTrendingMovies(ctx context.Context, input movie.GetTrendingMoviesInput) (movie.GetTrendingMoviesOutput, error) {
	movies, err := uc.repo.GetTrendingMovies(ctx, movie.GetTrendingMoviesOptions{
		Filter: input.Filter,
		Type:   input.Type,
	})
	if err != nil {
		return movie.GetTrendingMoviesOutput{}, err
	}

	return movie.GetTrendingMoviesOutput{Movies: movies}, nil
}

func (uc implUsecase) GetTopRatedMovies(ctx context.Context, input movie.GetTopRatedMoviesInput) (movie.GetTopRatedMoviesOutput, error) {
	movies, err := uc.repo.GetTopRatedMovies(ctx, movie.GetTopRatedMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetTopRatedMoviesOutput{}, err
	}

	return movie.GetTopRatedMoviesOutput{Movies: movies}, nil
}

func (uc implUsecase) GetPopularMovies(ctx context.Context, input movie.GetPopularMoviesInput) (movie.GetPopularMoviesOutput, error) {
	movies, err := uc.repo.GetPopularMovies(ctx, movie.GetPopularMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetPopularMoviesOutput{}, err
	}

	return movie.GetPopularMoviesOutput{Movies: movies}, nil
}

func (uc implUsecase) GetMovieGenres(ctx context.Context) (movie.GetMovieGenresOutput, error) {
	genres, err := uc.repo.GetMovieGenres(ctx)
	if err != nil {
		return movie.GetMovieGenresOutput{}, err
	}

	return movie.GetMovieGenresOutput{Genres: genres}, nil
}
