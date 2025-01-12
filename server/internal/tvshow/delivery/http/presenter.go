package http

import (
	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/tvshow"
	"github.com/tmplam/movseek/pkg"
)

type filter struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page" binding:"required"`
	// filter
	StartDate        string  `form:"start_date"`
	EndDate          string  `form:"end_date"`
	GenreIDs         []int64 `form:"genre_ids"`
	StartAverageVote float64 `form:"start_average_vote"`
	EndAverageVote   float64 `form:"end_average_vote"`
	// sort order
	TimeOrder       int `form:"time_order"`
	PopularityOrder int `form:"popularity"`
	VoteOrder       int `form:"vote_average"`
	NameOrder       int `form:"name"`
}

type searchTVShowsRequest struct {
	Query     string   `form:"query"`
	IDs       []int64  `form:"ids"`
	ObjectIDs []string `form:"object_ids"`
	Filter    filter   `form:"filter"`
}

func (req searchTVShowsRequest) toInput() tvshow.ListTVsInput {
	filter := tvshow.GetTVFilter{
		Page:             req.Filter.Page,
		PerPage:          req.Filter.PerPage,
		IDs:              req.IDs,
		ObjectIDs:        req.ObjectIDs,
		StartDate:        req.Filter.StartDate,
		EndDate:          req.Filter.EndDate,
		GenreIDs:         req.Filter.GenreIDs,
		StartAverageVote: pkg.CalculateScore(req.Filter.StartAverageVote),
		EndAverageVote:   pkg.CalculateScore(req.Filter.EndAverageVote),
		TimeOrder:        req.Filter.TimeOrder,
		PopularityOrder:  req.Filter.PopularityOrder,
		VoteOrder:        req.Filter.VoteOrder,
		NameOrder:        req.Filter.NameOrder,
	}

	return tvshow.ListTVsInput{
		Query:  req.Query,
		Filter: filter,
	}
}

type getUpcomingTVShowsRequest struct {
	filter
}

func (req getUpcomingTVShowsRequest) toInput() tvshow.GetUpcomingTVsInput {
	filter := tvshow.GetTVFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: pkg.CalculateScore(req.StartAverageVote),
		EndAverageVote:   pkg.CalculateScore(req.EndAverageVote),
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		NameOrder:        req.NameOrder,
	}

	return tvshow.GetUpcomingTVsInput{
		Filter: filter,
	}
}

type getOnTheAirTVShowsRequest struct {
	filter
}

func (req getOnTheAirTVShowsRequest) toInput() tvshow.GetOnTheAirTVsInput {
	filter := tvshow.GetTVFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: pkg.CalculateScore(req.StartAverageVote),
		EndAverageVote:   pkg.CalculateScore(req.EndAverageVote),
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		NameOrder:        req.NameOrder,
	}

	return tvshow.GetOnTheAirTVsInput{
		Filter: filter,
	}
}

type getAiringTodayTVShowsRequest struct {
	filter
}

func (req getAiringTodayTVShowsRequest) toInput() tvshow.GetAiringTodayTVsInput {
	filter := tvshow.GetTVFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: pkg.CalculateScore(req.StartAverageVote),
		EndAverageVote:   pkg.CalculateScore(req.EndAverageVote),
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		NameOrder:        req.NameOrder,
	}

	return tvshow.GetAiringTodayTVsInput{
		Filter: filter,
	}
}

type getTopRatedTVShowsRequest struct {
	filter
}

func (req getTopRatedTVShowsRequest) toInput() tvshow.GetTopRatedTVsInput {
	filter := tvshow.GetTVFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: pkg.CalculateScore(req.StartAverageVote),
		EndAverageVote:   pkg.CalculateScore(req.EndAverageVote),
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		NameOrder:        req.NameOrder,
	}

	return tvshow.GetTopRatedTVsInput{
		Filter: filter,
	}
}

type getPopularTVShowsRequest struct {
	filter `form:"filter"`
}

func (req getPopularTVShowsRequest) toInput() tvshow.GetPopularTVsInput {
	filter := tvshow.GetTVFilter{
		Page:             req.Page,
		PerPage:          req.PerPage,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		GenreIDs:         req.GenreIDs,
		StartAverageVote: pkg.CalculateScore(req.StartAverageVote),
		EndAverageVote:   pkg.CalculateScore(req.EndAverageVote),
		TimeOrder:        req.TimeOrder,
		PopularityOrder:  req.PopularityOrder,
		VoteOrder:        req.VoteOrder,
		NameOrder:        req.NameOrder,
	}

	return tvshow.GetPopularTVsInput{
		Filter: filter,
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
