package movie

import (
	"github.com/tmplam/movseek/internal/movie"
)

type implUsecase struct {
	repo movie.Repository
}

func NewUsecase(repo movie.Repository) movie.Usecase {
	return &implUsecase{
		repo: repo,
	}
}
