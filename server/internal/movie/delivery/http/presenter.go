package http

import (
	"errors"

	"github.com/tmplam/movseek/internal/movie"
)

type filter struct {
	page    int `form:"page"`
	perPage int `form:"per_page"`
}

type searchMoviesRequest struct {
	query string `form:"query"`
	filter
}

func (req searchMoviesRequest) toInput() movie.ListMoviesInput {
	return movie.ListMoviesInput{
		Query: req.query,
		Filter: movie.GetMovieFilter{
			Page:    req.page,
			PerPage: req.perPage,
		},
	}
}

type getUpcomingMoviesRequest struct {
	filter
}

func (req getUpcomingMoviesRequest) toInput() movie.GetUpcomingMoviesInput {
	return movie.GetUpcomingMoviesInput{
		Filter: movie.GetMovieFilter{
			Page:    req.page,
			PerPage: req.perPage,
		},
	}
}

type getTrendingMoviesRequest struct {
	t string `uri:"type"`
	filter
}

func (req getTrendingMoviesRequest) validate() error {
	if req.t != "week" && req.t != "day" {
		return errors.New("invalid type")
	}

	return nil
}

func (req getTrendingMoviesRequest) toInput() movie.GetTrendingMoviesInput {
	return movie.GetTrendingMoviesInput{
		Type: req.t,
		Filter: movie.GetMovieFilter{
			Page:    req.page,
			PerPage: req.perPage,
		},
	}
}

type getTopRatedMoviesRequest struct {
	filter
}

func (req getTopRatedMoviesRequest) toInput() movie.GetTopRatedMoviesInput {
	return movie.GetTopRatedMoviesInput{
		Filter: movie.GetMovieFilter{
			Page:    req.page,
			PerPage: req.perPage,
		},
	}
}

type getPopularMoviesRequest struct {
	filter
}

func (req getPopularMoviesRequest) toInput() movie.GetPopularMoviesInput {
	return movie.GetPopularMoviesInput{
		Filter: movie.GetMovieFilter{
			Page:    req.page,
			PerPage: req.perPage,
		},
	}
}
