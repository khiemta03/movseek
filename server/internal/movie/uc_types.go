package movie

import (
	"github.com/tmplam/movseek/internal/models"
)

type GetMovieFilter struct {
	Page             int     `json:"page"`
	PerPage          int     `json:"per_page"`
	StartDate        string  `json:"start_date"`
	EndDate          string  `json:"end_date"`
	GenreIDs         []int64 `json:"genre_ids"`
	StartAverageVote float64 `json:"start_average_vote"`
	EndAverageVote   float64 `json:"end_average_vote"`
	// sort order
	TimeOrder       int `json:"time_order"`
	PopularityOrder int `json:"popularity"`
	VoteOrder       int `json:"vote_average"`
	TitleOrder      int `json:"title"`

	IDs       []int64  `json:"ids"`
	ObjectIDs []string `json:"object_ids"`
}

type Pagination struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type ListMoviesInput struct {
	Query  string `json:"query"`
	Filter GetMovieFilter
}

type ListMoviesOutput struct {
	Movies     []models.Movie `json:"movies"`
	Pagination Pagination     `json:"pagination"`
}

type GetUpcomingMoviesInput struct {
	Filter GetMovieFilter
}

type GetUpcomingMoviesOutput struct {
	Movies     []models.MovieSummary `json:"movies"`
	Pagination Pagination            `json:"pagination"`
}

type GetTrendingMoviesInput struct {
	Filter GetMovieFilter
	Type   string
}

type GetTrendingMoviesOutput struct {
	Movies     []models.MovieSummary `json:"movies"`
	Pagination Pagination            `json:"pagination"`
}

type GetTopRatedMoviesInput struct {
	Filter GetMovieFilter
}

type GetTopRatedMoviesOutput struct {
	Movies     []models.MovieSummary `json:"movies"`
	Pagination Pagination            `json:"pagination"`
}

type GetPopularMoviesInput struct {
	Filter GetMovieFilter
}

type GetPopularMoviesOutput struct {
	Movies     []models.MovieSummary `json:"movies"`
	Pagination Pagination            `json:"pagination"`
}

type GetNowPlayingMoviesInput struct {
	Filter GetMovieFilter
}

type GetNowPlayingMoviesOutput struct {
	Movies     []models.MovieSummary `json:"movies"`
	Pagination Pagination            `json:"pagination"`
}

type GetMovieGenresOutput struct {
	Genres []models.MovieGenre `json:"genres"`
}

type GetLatestMoviesInput struct {
	Filter GetMovieFilter
}

type GetLastestTrailerInput struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
