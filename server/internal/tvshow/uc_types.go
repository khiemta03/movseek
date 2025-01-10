package tvshow

import "github.com/tmplam/movseek/internal/models"

type GetTVFilter struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type ListTVsInput struct {
	Query  string `json:"query"`
	Filter GetTVFilter
}

type ListTVsOutput struct {
	TVShows []models.TVShow `json:"TVShows"`
}

type GetUpcomingTVsInput struct {
	Filter GetTVFilter
}

type GetUpcomingTVsOutput struct {
	TVShows []models.TVSummary `json:"TVShows"`
}

type GetTrendingTVsInput struct {
	Filter GetTVFilter
	Type   string
}

type GetTrendingTVsOutput struct {
	TVShows []models.TVSummary `json:"TVShows"`
}

type GetTopRatedTVsInput struct {
	Filter GetTVFilter
}

type GetTopRatedTVsOutput struct {
	TVShows []models.TVSummary `json:"TVShows"`
}

type GetPopularTVsInput struct {
	Filter GetTVFilter
}

type GetPopularTVsOutput struct {
	TVShows []models.TVSummary `json:"TVShows"`
}

type GetTVGenresOutput struct {
	Genres []models.TVGenre `json:"genres"`
}
