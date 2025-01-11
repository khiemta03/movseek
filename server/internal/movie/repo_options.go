package movie

import "github.com/tmplam/movseek/internal/models"

type ListMoviesOptions struct {
	Query  string
	Filter GetMovieFilter
}

type GetUpcomingMoviesOptions struct {
	Filter GetMovieFilter
}

type GetTrendingMoviesOptions struct {
	Filter GetMovieFilter
	Type   string
}

type GetNowPlayingMoviesOptions struct {
	Filter GetMovieFilter
}

type GetTopRatedMoviesOptions struct {
	Filter GetMovieFilter
}

type GetPopularMoviesOptions struct {
	Filter GetMovieFilter
}

type GetLastestTrailerResponse struct {
	Trailer      models.Trailer `json:"trailer"`
	ID           int64          `json:"id"`
	PosterPath   string         `json:"poster_path"`
	BackdropPath string         `json:"backdrop_path"`
	Title        string         `json:"title"`
}
