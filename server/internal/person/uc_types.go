package person

import "github.com/tmplam/movseek/internal/models"

type GetPersonFilter struct {
	Page    int     `json:"page"`
	PerPage int     `json:"per_page"`
	IDs     []int64 `json:"ids"`
}

type Pagination struct {
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type ListPeopleInput struct {
	Query  string `json:"query"`
	Filter GetPersonFilter
}

type ListPeopleOutput struct {
	People     []models.Person `json:"people"`
	Pagination Pagination      `json:"pagination"`
}

type GetTrendingPeopleInput struct {
	Filter GetPersonFilter
	Type   string
}

type GetTrendingPeopleOutput struct {
	People     []models.PersonSummary `json:"people"`
	Pagination Pagination             `json:"pagination"`
}

type GetPopularPeopleInput struct {
	Filter GetPersonFilter
}

type GetPopularPeopleOutput struct {
	People     []models.PersonSummary `json:"people"`
	Pagination Pagination             `json:"pagination"`
}
