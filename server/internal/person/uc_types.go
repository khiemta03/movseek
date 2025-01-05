package person

import "github.com/tmplam/movseek/internal/models"

type GetPersonFilter struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type ListPeopleInput struct {
	Query  string `json:"query"`
	Filter GetPersonFilter
}

type ListPeopleOutput struct {
	People []models.Person `json:"people"`
}

type GetTrendingPeopleInput struct {
	Filter GetPersonFilter
	Type   string
}

type GetTrendingPeopleOutput struct {
	People []models.PersonSummary `json:"people"`
}

type GetPopularPeopleInput struct {
	Filter GetPersonFilter
}

type GetPopularPeopleOutput struct {
	People []models.PersonSummary `json:"people"`
}
