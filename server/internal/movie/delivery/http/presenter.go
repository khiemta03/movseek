package http

import (
	"errors"

	"github.com/tmplam/movseek/internal/movie"
)

type filter struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page" binding:"required"`
}

type searchMoviesRequest struct {
	Query  string `form:"query"`
	Filter filter `form:"filter"`
}

func (req searchMoviesRequest) toInput() movie.ListMoviesInput {
	return movie.ListMoviesInput{
		Query: req.Query,
		Filter: movie.GetMovieFilter{
			Page:    req.Filter.Page,
			PerPage: req.Filter.PerPage,
		},
	}
}

type getUpcomingMoviesRequest struct {
	filter
}

func (req getUpcomingMoviesRequest) toInput() movie.GetUpcomingMoviesInput {
	return movie.GetUpcomingMoviesInput{
		Filter: movie.GetMovieFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getTrendingMoviesRequest struct {
	Type string `uri:"type"`
	filter
}

func (req getTrendingMoviesRequest) validate() error {
	if req.Type != "week" && req.Type != "day" {
		return errors.New("invalid type")
	}

	return nil
}

func (req getTrendingMoviesRequest) toInput() movie.GetTrendingMoviesInput {
	return movie.GetTrendingMoviesInput{
		Type: req.Type,
		Filter: movie.GetMovieFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getTopRatedMoviesRequest struct {
	filter
}

func (req getTopRatedMoviesRequest) toInput() movie.GetTopRatedMoviesInput {
	return movie.GetTopRatedMoviesInput{
		Filter: movie.GetMovieFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getPopularMoviesRequest struct {
	filter `form:"filter"`
}

func (req getPopularMoviesRequest) toInput() movie.GetPopularMoviesInput {
	return movie.GetPopularMoviesInput{
		Filter: movie.GetMovieFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}
