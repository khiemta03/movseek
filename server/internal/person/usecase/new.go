package movie

import (
	"github.com/tmplam/movseek/internal/person"
)

type implUsecase struct {
	repo person.Repository
}

func NewUsecase(repo person.Repository) person.Usecase {
	return &implUsecase{
		repo: repo,
	}
}
