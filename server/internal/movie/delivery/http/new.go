package http

import "github.com/tmplam/movseek/internal/movie"

type handler struct {
	uc movie.Usecase
}

func NewHandler(uc movie.Usecase) *handler {
	return &handler{uc: uc}
}
