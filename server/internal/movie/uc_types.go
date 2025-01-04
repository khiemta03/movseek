package movie

import "github.com/tmplam/movseek/internal/models"

type ListMoviesInput struct {
	Query   string `json:"query"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

type ListMoviesOutput struct {
	Movies []models.Movie `json:"movies"`
}
