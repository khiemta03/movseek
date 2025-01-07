package tvshow

type ListTVsOptions struct {
	Query  string
	Filter GetTVFilter
}

type GetUpcomingTVsOptions struct {
	Filter GetTVFilter
}

type GetTrendingTVsOptions struct {
	Filter GetTVFilter
	Type   string
}

type GetTopRatedTVsOptions struct {
	Filter GetTVFilter
}

type GetPopularTVsOptions struct {
	Filter GetTVFilter
}
