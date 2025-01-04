package http

import (
	"github.com/tmplam/movseek/internal/movie"
)

type searchMoviesRequest struct {
	Query   string `form:"query"`
	Page    int    `form:"page"`
	PerPage int    `form:"per_page"`
}

func (req searchMoviesRequest) toInput() movie.ListMoviesInput {
	return movie.ListMoviesInput{
		Query:   req.Query,
		Page:    req.Page,
		PerPage: req.PerPage,
	}
}
