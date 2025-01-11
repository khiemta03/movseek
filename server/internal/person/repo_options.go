package person

type ListPeopleOptions struct {
	Query  string
	Filter GetPersonFilter
}

type GetTrendingPeopleOptions struct {
	Filter GetPersonFilter
	Type   string
}

type GetPopularPeopleOptions struct {
	Filter GetPersonFilter
}
