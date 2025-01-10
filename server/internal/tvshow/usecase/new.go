package movie

import "github.com/tmplam/movseek/internal/tvshow"

type implUsecase struct {
	repo tvshow.Repository
}

func NewUsecase(repo tvshow.Repository) tvshow.Usecase {
	return &implUsecase{
		repo: repo,
	}
}
