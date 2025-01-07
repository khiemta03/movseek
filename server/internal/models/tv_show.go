package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TVShow struct {
	ID                  primitive.ObjectID    `json:"_id" bson:"_id"`
	TMDBID              int64                 `json:"tmdb_id" bson:"tmdb_id"`
	Adult               bool                  `json:"adult" bson:"adult"`
	BackdropPath        string                `json:"backdrop_path" bson:"backdrop_path"`
	CreatedBy           []CreatedBy           `json:"created_by" bson:"created_by"`
	EpisodeRunTime      []int64               `json:"episode_run_time" bson:"episode_run_time"`
	FirstAirDate        string                `json:"first_air_date" bson:"first_air_date"`
	Genres              []TVShowGenres        `json:"genres" bson:"genres"`
	Homepage            string                `json:"homepage" bson:"homepage"`
	BelongsTo           BelongsTo             `json:"belongs_to_collection" bson:"belongs_to_collection"`
	InProduction        bool                  `json:"in_production" bson:"in_production"`
	Languages           []string              `json:"languages" bson:"languages"`
	LastAirDate         string                `json:"last_air_date" bson:"last_air_date"`
	LastEpisodeToAir    EpisodeToAir          `json:"last_episode_to_air" bson:"last_episode_to_air"`
	Name                string                `json:"name" bson:"name"`
	Networks            []Networks            `json:"networks" bson:"networks"`
	NextEpisodeToAir    EpisodeToAir          `json:"next_episode_to_air" bson:"next_episode_to_air"`
	NumberEpisodes      int64                 `json:"number_of_episodes" bson:"number_of_episodes"`
	NumberSeasons       int64                 `json:"number_of_seasons" bson:"number_of_seasons"`
	OriginCountry       []string              `json:"origin_country" bson:"origin_country"`
	OriginalLanguage    string                `json:"original_language" bson:"original_language"`
	OriginalName        string                `json:"original_name" bson:"original_name"`
	Overview            string                `json:"overview" bson:"overview"`
	Popularity          float64               `json:"popularity" bson:"popularity"`
	PosterPath          string                `json:"poster_path" bson:"poster_path"`
	ProductionCompanies []ProductionCompanies `json:"production_companies" bson:"production_companies"`
	ProductionCountries []ProductionCountries `json:"production_countries" bson:"production_countries"`
	SpokenLanguages     []SpokenLanguages     `json:"spoken_languages" bson:"spoken_languages"`
	Status              string                `json:"status" bson:"status"`
	Tagline             string                `json:"tagline" bson:"tagline"`
	Type                string                `json:"type" bson:"type"`
	VoteAverage         float64               `json:"vote_average" bson:"vote_average"`
	VoteCount           int64                 `json:"vote_count" bson:"vote_count"`
}

type CreatedBy struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	ProfilePath  string             `json:"profile_path" bson:"profile_path"`
	CreditID     string             `json:"credit_id" bson:"credit_id"`
	OriginalName string             `json:"original_name" bson:"original_name"`
	Gender       int                `json:"gender" bson:"gender"`
}

type TVShowGenres struct {
	ID   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type EpisodeToAir struct {
	ID             int64   `json:"id" bson:"id"`
	Name           string  `json:"name" bson:"name"`
	Overview       string  `json:"overview" bson:"overview"`
	VoteAverage    float64 `json:"vote_average" bson:"vote_average"`
	VoteCount      int64   `json:"vote_count" bson:"vote_count"`
	AirDate        string  `json:"air_date" bson:"air_date"`
	EpisodeNumber  int64   `json:"episode_number" bson:"episode_number"`
	EpisodeType    string  `json:"episode_type" bson:"episode_type"`
	ProductionCode string  `json:"production_code" bson:"production_code"`
	Runtime        int64   `json:"runtime" bson:"runtime"`
	SeasonNumber   int64   `json:"season_number" bson:"season_number"`
	ShowID         int64   `json:"show_id" bson:"show_id"`
	StillPath      string  `json:"still_path" bson:"still_path"`
}

type Networks struct {
	ID            int64  `json:"id" bson:"id"`
	Name          string `json:"name" bson:"name"`
	LogoPath      string `json:"logo_path" bson:"logo_path"`
	OriginCountry string `json:"origin_country" bson:"origin_country"`
}

type TVSummary struct {
	MovieID          primitive.ObjectID `json:"_id" bson:"_id"`
	TMDBID           int64              `json:"tmdb_id" bson:"tmdb_id"`
	Adult            bool               `json:"adult" bson:"adult"`
	BackdropPath     string             `json:"backdrop_path" bson:"backdrop_path"`
	GenreIDs         []int64            `json:"genre_ids" bson:"genre_ids"`
	OriginalLanguage string             `json:"original_language" bson:"original_language"`
	OriginalTitle    string             `json:"original_title" bson:"original_title"`
	Overview         string             `json:"overview" bson:"overview"`
	Popularity       float64            `json:"popularity" bson:"popularity"`
	PosterPath       string             `json:"poster_path" bson:"poster_path"`
	ReleaseDate      string             `json:"release_date" bson:"release_date"`
	Title            string             `json:"title" bson:"title"`
	Video            bool               `json:"video" bson:"video"`
	VoteAverage      float64            `json:"vote_average" bson:"vote_average"`
	VoteCount        int64              `json:"vote_count" bson:"vote_count"`
}
