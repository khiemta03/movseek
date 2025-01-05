package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	MovieID          primitive.ObjectID `json:"_id" bson:"_id"`
	TMDBID           int64              `json:"tmdb_id" bson:"tmdb_id"`
	BackdropPath     string             `json:"backdrop_path" bson:"backdrop_path"`
	Categories       []string           `json:"categories" bson:"categories"`
	Genres           []MovieGenres      `json:"genres" bson:"genres"`
	ID               int64              `json:"id" bson:"id"`
	IMDBID           string             `json:"imdb_id" bson:"imdb_id"`
	OriginalCountry  []string           `json:"original_country" bson:"original_country"`
	OriginalLanguage string             `json:"original_language" bson:"original_language"`
	Overview         string             `json:"overview" bson:"overview"`
	Popularity       float64            `json:"popularity" bson:"popularity"`
	PosterPath       string             `json:"poster_path" bson:"poster_path"`
	ReleaseDate      string             `json:"release_date" bson:"release_date"`
	Revenue          int64              `json:"revenue" bson:"revenue"`
	Runtime          int64              `json:"runtime" bson:"runtime"`
	Status           string             `json:"status" bson:"status"`
	Tagline          string             `json:"tagline" bson:"tagline"`
	Title            string             `json:"title" bson:"title"`
	Video            bool               `json:"video" bson:"video"`
	VoteAverage      float64            `json:"vote_average" bson:"vote_average"`
	VoteCount        int64              `json:"vote_count" bson:"vote_count"`
	Credits          MovieCredits       `json:"credits" bson:"credits"`
}

type MovieGenres struct {
	ID   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type MovieCredits struct {
	ID   int64       `json:"id" bson:"id"`
	Cast []MovieCrew `json:"cast" bson:"cast"`
	Crew []MovieCrew `json:"crew" bson:"crew"`
}

type MovieCrew struct {
	Adult              bool    `json:"adult" bson:"adult"`
	Gender             int64   `json:"gender" bson:"gender"`
	ID                 int64   `json:"id" bson:"id"`
	KnownForDepartment string  `json:"known_for_department" bson:"known_for_department"`
	Name               string  `json:"name" bson:"name"`
	OriginalName       string  `json:"original_name" bson:"original_name"`
	Popularity         float64 `json:"popularity" bson:"popularity"`
	ProfilePath        string  `json:"profile_path" bson:"profile_path"`
	CastID             int64   `json:"cast_id" bson:"cast_id"`
	Character          string  `json:"character" bson:"character"`
	CreditID           string  `json:"credit_id" bson:"credit_id"`
	Order              int64   `json:"order" bson:"order"`
}

type MovieSummary struct {
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
