package http

import (
	"errors"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie"
)

type filter struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page" binding:"required"`
}

type pagination struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type searchMoviesRequest struct {
	Query     string   `form:"query"`
	Filter    filter   `form:"filter"`
	IDs       []int64  `form:"ids"`
	ObjectIDs []string `form:"object_ids"`
}

type searchMoviesResponse struct {
	Movies     []models.Movie `json:"movies"`
	Pagination pagination     `json:"pagination"`
}

func (req searchMoviesRequest) toInput() movie.ListMoviesInput {
	return movie.ListMoviesInput{
		Query: req.Query,
		Filter: movie.GetMovieFilter{
			Page:      req.Filter.Page,
			PerPage:   req.Filter.PerPage,
			IDs:       req.IDs,
			ObjectIDs: req.ObjectIDs,
		},
	}
}

func (h handlerImpl) buildSearchMoviesResponse(resp movie.ListMoviesOutput) searchMoviesResponse {
	return searchMoviesResponse{
		Movies: resp.Movies,
		Pagination: pagination{
			Page:      resp.Pagination.Page,
			PerPage:   resp.Pagination.PerPage,
			Total:     resp.Pagination.Total,
			TotalPage: resp.Pagination.TotalPage,
		},
	}
}

type getSummaryMoviesResponse struct {
	Movies     []models.MovieSummary `json:"movies"`
	Pagination pagination            `json:"pagination"`
}

func (h handlerImpl) buildGetSummaryMoviesResponse(movies []models.MovieSummary, pag movie.Pagination) getSummaryMoviesResponse {
	return getSummaryMoviesResponse{
		Movies: movies,
		Pagination: pagination{
			Page:      pag.Page,
			PerPage:   pag.PerPage,
			Total:     pag.Total,
			TotalPage: pag.TotalPage,
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

type getNowPlayingMoviesRequest struct {
	filter `form:"filter"`
}

func (req getNowPlayingMoviesRequest) toInput() movie.GetNowPlayingMoviesInput {
	return movie.GetNowPlayingMoviesInput{
		Filter: movie.GetMovieFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getLastestTrailerRequest struct {
	filter
}

func (req getLastestTrailerRequest) toInput() movie.GetLastestTrailerInput {
	return movie.GetLastestTrailerInput{
		Page:    req.Page,
		PerPage: req.PerPage,
	}
}
