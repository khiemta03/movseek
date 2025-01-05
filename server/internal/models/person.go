package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id"`
	TMDBID            int                `json:"tmdb_id" bson:"tmdb_id"`
	Adult             bool               `json:"adult" bson:"adult"`
	Gender            int                `json:"gender" bson:"gender"`
	AlsoKnowAs        []string           `json:"also_known_as" bson:"also_known_as"`
	Biography         string             `json:"biography" bson:"biography"`
	Birthday          string             `json:"birthday" bson:"birthday"`
	Deathday          *string            `json:"deathday" bson:"deathday"`
	Homepage          string             `json:"homepage" bson:"homepage"`
	ImdbID            string             `json:"imdb_id" bson:"imdb_id"`
	PlaceOfBirth      string             `json:"place_of_birth" bson:"place_of_birth"`
	MovieCredits      MovieCredit        `json:"movie_credits" bson:"movie_credits"`
	Name              string             `json:"name" bson:"name"`
	KnowForDepartment string             `json:"know_for_department" bson:"know_for_department"`
	OriginalName      string             `json:"original_name" bson:"original_name"`
	Popularity        float64            `json:"popularity" bson:"popularity"`
	ProfilePath       string             `json:"profile_path" bson:"profile_path"`
}

type KnownFor struct {
	ID               int     `json:"id" bson:"id"`
	BackdropPath     string  `json:"backdrop_path" bson:"backdrop_path"`
	PosterPath       string  `json:"poster_path" bson:"poster_path"`
	Title            string  `json:"title" bson:"title"`
	Overview         string  `json:"overview" bson:"overview"`
	OriginalTitle    string  `json:"original_title" bson:"original_title"`
	MediaType        string  `json:"media_type" bson:"media_type"`
	Adult            bool    `json:"adult" bson:"adult"`
	OriginalLanguage string  `json:"original_language" bson:"original_language"`
	GenreIDs         []int   `json:"genre_ids" bson:"genre_ids"`
	Popularity       float64 `json:"popularity" bson:"popularity"`
	ReleaseDate      string  `json:"release_date" bson:"release_date"`
	Video            bool    `json:"video" bson:"video"`
	VoteAverage      float64 `json:"vote_average" bson:"vote_average"`
	VoteCount        int     `json:"vote_count" bson:"vote_count"`
}

type PersonSummary struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id"`
	TMDBID            int                `json:"tmdb_id" bson:"tmdb_id"`
	Adult             bool               `json:"adult" bson:"adult"`
	Gender            int                `json:"gender" bson:"gender"`
	KnownFor          []KnownFor         `json:"known_for" bson:"known_for"`
	Name              string             `json:"name" bson:"name"`
	KnowForDepartment string             `json:"know_for_department" bson:"know_for_department"`
	OriginalName      string             `json:"original_name" bson:"original_name"`
	Popularity        float64            `json:"popularity" bson:"popularity"`
	ProfilePath       string             `json:"profile_path" bson:"profile_path"`
}

type MovieCredit struct {
	ID   int64  `json:"id" bson:"id"`
	Cast []Crew `json:"cast" bson:"cast"`
	Crew []Crew `json:"crew" bson:"crew"`
}

type Crew struct {
	ID               int64   `json:"id" bson:"id"`
	Adult            bool    `json:"adult" bson:"adult"`
	BackdropPath     *string `json:"backdrop_path" bson:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids" bson:"genre_ids"`
	PosterPath       string  `json:"poster_path" bson:"poster_path"`
	OriginalTitle    string  `json:"original_title" bson:"original_title"`
	OriginalLanguage string  `json:"original_language" bson:"original_language"`
	Title            string  `json:"title" bson:"title"`
	Overview         string  `json:"overview" bson:"overview"`
	ReleaseDate      string  `json:"release_date" bson:"release_date"`
	Video            bool    `json:"video" bson:"video"`
	VoteAverage      float64 `json:"vote_average" bson:"vote_average"`
	VoteCount        int     `json:"vote_count" bson:"vote_count"`
	Character        string  `json:"character" bson:"character"`
	CreditID         string  `json:"credit_id" bson:"credit_id"`
	Order            int64   `json:"order" bson:"order"`
}
