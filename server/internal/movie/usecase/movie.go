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
		Query:   input.Query,
		Page:    input.Page,
		PerPage: input.PerPage,
	})
	if err != nil {
		return movie.ListMoviesOutput{}, err
	}

	return movie.ListMoviesOutput{Movies: movies}, nil
}
