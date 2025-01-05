package movie

import "github.com/tmplam/movseek/internal/models"

type GetMovieFilter struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type ListMoviesInput struct {
	Query  string `json:"query"`
	Filter GetMovieFilter
}

type ListMoviesOutput struct {
	Movies []models.Movie `json:"movies"`
}

type GetUpcomingMoviesInput struct {
	Filter GetMovieFilter
}

type GetUpcomingMoviesOutput struct {
	Movies []models.MovieSummary `json:"movies"`
}

type GetTrendingMoviesInput struct {
	Filter GetMovieFilter
	Type   string
}

type GetTrendingMoviesOutput struct {
	Movies []models.MovieSummary `json:"movies"`
}
