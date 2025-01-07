package http

import (
	"errors"

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

type getTrendingTVShowsRequest struct {
	Type string `uri:"type"`
	filter
}

func (req getTrendingTVShowsRequest) validate() error {
	if req.Type != "week" && req.Type != "day" {
		return errors.New("invalid type")
	}

	return nil
}

func (req getTrendingTVShowsRequest) toInput() tvshow.GetTrendingTVsInput {
	return tvshow.GetTrendingTVsInput{
		Type: req.Type,
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
