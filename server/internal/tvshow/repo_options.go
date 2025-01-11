package tvshow

type ListTVsOptions struct {
	Query  string
	Filter GetTVFilter
}

type GetUpcomingTVsOptions struct {
	Filter GetTVFilter
}

type GetOnTheAirTVsOptions struct {
	Filter GetTVFilter
}

type GetAiringTodayTVsOptions struct {
	Filter GetTVFilter
}

type GetTopRatedTVsOptions struct {
	Filter GetTVFilter
}

type GetPopularTVsOptions struct {
	Filter GetTVFilter
}
