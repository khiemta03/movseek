package http

import (
	"errors"
	"fmt"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/person"
)

type filter struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page" binding:"required"`
}

type searchPeopleRequest struct {
	Query string `form:"query"`
	filter
}

func (req searchPeopleRequest) toInput() person.ListPeopleInput {
	fmt.Println("query", req.Query)
	return person.ListPeopleInput{
		Query: req.Query,
		Filter: person.GetPersonFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getTrendingPeopleRequest struct {
	Type string `uri:"type"`
	filter
}

func (req getTrendingPeopleRequest) validate() error {
	if req.Type != "week" && req.Type != "day" {
		return errors.New("invalid type")
	}

	return nil
}

func (req getTrendingPeopleRequest) toInput() person.GetTrendingPeopleInput {
	return person.GetTrendingPeopleInput{
		Type: req.Type,
		Filter: person.GetPersonFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getPopularPeopleRequest struct {
	filter
}

func (req getPopularPeopleRequest) toInput() person.GetPopularPeopleInput {
	return person.GetPopularPeopleInput{
		Filter: person.GetPersonFilter{
			Page:    req.Page,
			PerPage: req.PerPage,
		},
	}
}

type getListPeopleResponse struct {
	Results      []models.Person `json:"results"`
	Page         int             `json:"page"`
	PerPage      int             `json:"per_page"`
	TotalPages   int             `json:"total_pages"`
	TotalResults int             `json:"total_results"`
}

func (h handlerImpl) toResponse(resp person.ListPeopleOutput) getListPeopleResponse {
	return getListPeopleResponse{
		Results:      resp.People,
		Page:         resp.Pagination.Page,
		PerPage:      resp.Pagination.PerPage,
		TotalPages:   resp.Pagination.TotalPages,
		TotalResults: resp.Pagination.TotalResults,
	}
}

type getListPeopleSummaryResponse struct {
	Results      []models.PersonSummary `json:"results"`
	Page         int                    `json:"page"`
	PerPage      int                    `json:"per_page"`
	TotalPages   int                    `json:"total_pages"`
	TotalResults int                    `json:"total_results"`
}

func (h handlerImpl) toSummaryResponse(people []models.PersonSummary, pagination person.Pagination) getListPeopleSummaryResponse {
	return getListPeopleSummaryResponse{
		Results:      people,
		Page:         pagination.Page,
		PerPage:      pagination.PerPage,
		TotalPages:   pagination.TotalPages,
		TotalResults: pagination.TotalResults,
	}
}
