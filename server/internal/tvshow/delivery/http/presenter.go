package http

import (
	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/tvshow"
)

type filter struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page" binding:"required"`
}

type searchTVShowsRequest struct {
	Query  string `form:"query"`
	Filter filter `form:"filter"`
}

func (req searchTVShowsRequest) toInput() tvshow.ListTVsInput {
	return tvshow.ListTVsInput{
		Query: req.Query,
		Filter: tvshow.GetTVFilter{
			Page:    req.Filter.Page,
			PerPage: req.Filter.PerPage,
		},
	}
}

type getUpcomingTVShowsRequest struct {
	filter
}

func (req getUpcomingTVShowsRequest) toInput() tvshow.GetUpcomingTVsInput {
	return tvshow.GetUpcomingTVsInput{
		Filter: tvshow.GetTVFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getOnTheAirTVShowsRequest struct {
	filter
}

func (req getOnTheAirTVShowsRequest) toInput() tvshow.GetOnTheAirTVsInput {
	return tvshow.GetOnTheAirTVsInput{
		Filter: tvshow.GetTVFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getAiringTodayTVShowsRequest struct {
	filter
}

func (req getAiringTodayTVShowsRequest) toInput() tvshow.GetAiringTodayTVsInput {
	return tvshow.GetAiringTodayTVsInput{
		Filter: tvshow.GetTVFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getTopRatedTVShowsRequest struct {
	filter
}

func (req getTopRatedTVShowsRequest) toInput() tvshow.GetTopRatedTVsInput {
	return tvshow.GetTopRatedTVsInput{
		Filter: tvshow.GetTVFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getPopularTVShowsRequest struct {
	filter `form:"filter"`
}

func (req getPopularTVShowsRequest) toInput() tvshow.GetPopularTVsInput {
	return tvshow.GetPopularTVsInput{
		Filter: tvshow.GetTVFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getListTVShowsResponse struct {
	Results      []models.TVShow `json:"results"`
	Page         int             `json:"page"`
	PerPage      int             `json:"per_page"`
	TotalResults int             `json:"total_results"`
	TotalPages   int             `json:"total_pages"`
}

func (h handlerImpl) toResponse(res tvshow.ListTVsOutput) getListTVShowsResponse {
	return getListTVShowsResponse{
		Results:      res.TVShows,
		Page:         res.Pagination.Page,
		PerPage:      res.Pagination.PerPage,
		TotalPages:   res.Pagination.TotalPages,
		TotalResults: res.Pagination.Total,
	}
}

type getListTVShowSummaryResponse struct {
	Results      []models.TVSummary `json:"results"`
	Page         int                `json:"page"`
	PerPage      int                `json:"per_page"`
	TotalResults int                `json:"total_results"`
	TotalPages   int                `json:"total_pages"`
}

func (h handlerImpl) toSummaryResponse(tvshows []models.TVSummary, pagination tvshow.Pagination) getListTVShowSummaryResponse {
	return getListTVShowSummaryResponse{
		Results:      tvshows,
		Page:         pagination.Page,
		PerPage:      pagination.PerPage,
		TotalPages:   pagination.TotalPages,
		TotalResults: pagination.Total,
	}
}
