package movie

import (
	"github.com/tmplam/movseek/internal/review"
)

type implUsecase struct {
	repo review.Repository
}

func NewUsecase(repo review.Repository) review.Usecase {
	return &implUsecase{
		repo: repo,
	}
}
