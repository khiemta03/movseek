package movie

import (
	"github.com/tmplam/movseek/internal/movie"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"github.com/tmplam/movseek/internal/tvshow"
)

type implUsecase struct {
	repo    saved_item.Repository
	movieUC movie.Usecase
	tvUC    tvshow.Usecase
}

func NewUsecase(repo saved_item.Repository, movieUC movie.Usecase, tvUC tvshow.Usecase) saved_item.Usecase {
	return &implUsecase{
		repo:    repo,
		movieUC: movieUC,
		tvUC:    tvUC,
	}
}
