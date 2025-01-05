package http

import (
	"errors"

	"github.com/tmplam/movseek/internal/person"
)

type filter struct {
	Page    int `form:"page" binding:"required"`
	PerPage int `form:"per_page" binding:"required"`
}

type searchPeopleRequest struct {
	query string `form:"query"`
	filter
}

func (req searchPeopleRequest) toInput() person.ListPeopleInput {
	return person.ListPeopleInput{
		Query: req.query,
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
