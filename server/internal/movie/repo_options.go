package movie

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
