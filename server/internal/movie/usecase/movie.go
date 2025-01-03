package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
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
