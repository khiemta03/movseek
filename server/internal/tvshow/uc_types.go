package tvshow

import "github.com/tmplam/movseek/internal/models"

type GetTVFilter struct {
	Page      int      `json:"page"`
	PerPage   int      `json:"per_page"`
	IDs       []int64  `json:"ids"`
	ObjectIDs []string `json:"object_ids"`

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
