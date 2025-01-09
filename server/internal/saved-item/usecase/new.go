package movie

import (
	saved_item "github.com/tmplam/movseek/internal/saved-item"
)

type implUsecase struct {
	repo saved_item.Repository
}

func NewUsecase(repo saved_item.Repository) saved_item.Usecase {
	return &implUsecase{
		repo: repo,
	}
}
