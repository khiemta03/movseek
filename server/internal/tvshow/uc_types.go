package tvshow

import "github.com/tmplam/movseek/internal/models"

type GetTVFilter struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ListTVsInput struct {
	Query  string `json:"query"`
	Filter GetTVFilter
}

type ListTVsOutput struct {
	TVShows    []models.TVShow `json:"TVShows"`
	Pagination Pagination      `json:"pagination"`
}

type GetUpcomingTVsInput struct {
	Filter GetTVFilter
}

type GetUpcomingTVsOutput struct {
	TVShows    []models.TVSummary `json:"TVShows"`
	Pagination Pagination         `json:"pagination"`
}

type GetOnTheAirTVsInput struct {
	Filter GetTVFilter
}

type GetOnTheAirTVsOutput struct {
	TVShows    []models.TVSummary `json:"TVShows"`
	Pagination Pagination         `json:"pagination"`
}

type GetAiringTodayTVsInput struct {
	Filter GetTVFilter
}

type GetAiringTodayTVsOutput struct {
	TVShows    []models.TVSummary `json:"TVShows"`
	Pagination Pagination         `json:"pagination"`
}

type GetTopRatedTVsInput struct {
	Filter GetTVFilter
}

type GetTopRatedTVsOutput struct {
	TVShows    []models.TVSummary `json:"TVShows"`
	Pagination Pagination         `json:"pagination"`
}

type GetPopularTVsInput struct {
	Filter GetTVFilter
}

type GetPopularTVsOutput struct {
	TVShows    []models.TVSummary `json:"TVShows"`
	Pagination Pagination         `json:"pagination"`
}

type GetTVGenresOutput struct {
	Genres []models.TVGenre `json:"genres"`
}
