package http

import (
	"errors"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie"
)

type filter struct {
	Page             int     `form:"page" binding:"required"`
	PerPage          int     `form:"per_page" binding:"required"`
	StartDate        string  `form:"start_date"`
	EndDate          string  `form:"end_date"`
	GenreIDs         []int64 `form:"genre_ids"`
	StartAverageVote float64 `form:"start_average_vote"`
	EndAverageVote   float64 `form:"end_average_vote"`
	// sort order
	TimeOrder       int `form:"time_order"`
	PopularityOrder int `form:"popularity"`
	VoteOrder       int `form:"vote_average"`
	TitleOrder      int `form:"title"`
}

type searchMoviesRequest struct {
	Query     string   `form:"query"`
	Filter    filter   `form:"filter"`
	IDs       []int64  `form:"ids"`
	ObjectIDs []string `form:"object_ids"`
}

type searchMoviesResponse struct {
	Results    []models.Movie `json:"results"`
	Page       int            `json:"page"`
	PerPage    int            `json:"per_page"`
	Total      int            `json:"total_results"`
	TotalPages int            `json:"total_pages"`
}

func (req searchMoviesRequest) toInput() movie.ListMoviesInput {
	filter := movie.GetMovieFilter{
		Page:             req.Filter.Page,
		PerPage:          req.Filter.PerPage,
		IDs:              req.IDs,
		ObjectIDs:        req.ObjectIDs,
		StartDate:        req.Filter.StartDate,
		EndDate:          req.Filter.EndDate,
		GenreIDs:         req.Filter.GenreIDs,
		StartAverageVote: req.Filter.StartAverageVote,
		EndAverageVote:   req.Filter.EndAverageVote,
		TimeOrder:        req.Filter.TimeOrder,
		PopularityOrder:  req.Filter.PopularityOrder,
		VoteOrder:        req.Filter.VoteOrder,
		TitleOrder:       req.Filter.TitleOrder,
	}

	return movie.ListMoviesInput{
		Query:  req.Query,
		Filter: filter,
	}
}

func (h handlerImpl) buildSearchMoviesResponse(resp movie.ListMoviesOutput) searchMoviesResponse {
	return searchMoviesResponse{
		Results:    resp.Movies,
		Page:       resp.Pagination.Page,
		PerPage:    resp.Pagination.PerPage,
		Total:      resp.Pagination.Total,
		TotalPages: resp.Pagination.TotalPage,
	}
}

type getSummaryMoviesResponse struct {
	Results    []models.MovieSummary `json:"results"`
	Page       int                   `json:"page"`
	PerPage    int                   `json:"per_page"`
	Total      int                   `json:"total_results"`
	TotalPages int                   `json:"total_pages"`
}

func (h handlerImpl) buildGetSummaryMoviesResponse(movies []models.MovieSummary, pag movie.Pagination) getSummaryMoviesResponse {
	return getSummaryMoviesResponse{
		Results:    movies,
		Page:       pag.Page,
		PerPage:    pag.PerPage,
		Total:      pag.Total,
		TotalPages: pag.TotalPage,
	}
}

type getUpcomingMoviesRequest struct {
	filter
}

func (req getUpcomingMoviesRequest) toInput() movie.GetUpcomingMoviesInput {
	filter := movie.GetMovieFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: req.StartAverageVote,
		EndAverageVote:   req.EndAverageVote,
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		TitleOrder:       req.TitleOrder,
	}

	return movie.GetUpcomingMoviesInput{
		Filter: filter,
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
	filter := movie.GetMovieFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: req.StartAverageVote,
		EndAverageVote:   req.EndAverageVote,
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		TitleOrder:       req.TitleOrder,
	}

	return movie.GetTrendingMoviesInput{
		Type:   req.Type,
		Filter: filter,
	}
}

type getTopRatedMoviesRequest struct {
	filter
}

func (req getTopRatedMoviesRequest) toInput() movie.GetTopRatedMoviesInput {
	filter := movie.GetMovieFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: req.StartAverageVote,
		EndAverageVote:   req.EndAverageVote,
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		TitleOrder:       req.TitleOrder,
	}

	return movie.GetTopRatedMoviesInput{
		Filter: filter,
	}
}

type getPopularMoviesRequest struct {
	filter `form:"filter"`
}

func (req getPopularMoviesRequest) toInput() movie.GetPopularMoviesInput {
	filter := movie.GetMovieFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: req.StartAverageVote,
		EndAverageVote:   req.EndAverageVote,
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		TitleOrder:       req.TitleOrder,
	}

	return movie.GetPopularMoviesInput{
		Filter: filter,
	}
}

type getNowPlayingMoviesRequest struct {
	filter `form:"filter"`
}

func (req getNowPlayingMoviesRequest) toInput() movie.GetNowPlayingMoviesInput {
	filter := movie.GetMovieFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: req.StartAverageVote,
		EndAverageVote:   req.EndAverageVote,
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		TitleOrder:       req.TitleOrder,
	}

	return movie.GetNowPlayingMoviesInput{
		Filter: filter,
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
