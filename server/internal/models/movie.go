package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	MovieID             primitive.ObjectID    `json:"_id" bson:"_id"`
	TMDBID              int64                 `json:"tmdb_id" bson:"tmdb_id"`
	BackdropPath        string                `json:"backdrop_path" bson:"backdrop_path"`
	BelongsTo           BelongsTo             `json:"belongs_to_collection" bson:"belongs_to_collection"`
	Budget              int64                 `json:"budget" bson:"budget"`
	Categories          []string              `json:"categories" bson:"categories"`
	Homepage            string                `json:"homepage" bson:"homepage"`
	Genres              []MovieGenres         `json:"genres" bson:"genres"`
	ID                  int64                 `json:"id" bson:"id"`
	IMDBID              string                `json:"imdb_id" bson:"imdb_id"`
	OriginalCountry     []string              `json:"original_country" bson:"original_country"`
	OriginalLanguage    string                `json:"original_language" bson:"original_language"`
	Overview            string                `json:"overview" bson:"overview"`
	Popularity          float64               `json:"popularity" bson:"popularity"`
	PosterPath          string                `json:"poster_path" bson:"poster_path"`
	ProductionCompanies []ProductionCompanies `json:"production_companies" bson:"production_companies"`
	ProductionCountries []ProductionCountries `json:"production_countries" bson:"production_countries"`
	SpokenLanguages     []SpokenLanguages     `json:"spoken_languages" bson:"spoken_languages"`
	ReleaseDate         string                `json:"release_date" bson:"release_date"`
	Revenue             int64                 `json:"revenue" bson:"revenue"`
	Runtime             int64                 `json:"runtime" bson:"runtime"`
	Status              string                `json:"status" bson:"status"`
	Tagline             string                `json:"tagline" bson:"tagline"`
	Title               string                `json:"title" bson:"title"`
	Video               bool                  `json:"video" bson:"video"`
	VoteAverage         float64               `json:"vote_average" bson:"vote_average"`
	VoteCount           int64                 `json:"vote_count" bson:"vote_count"`
	Credits             MovieCredits          `json:"credits" bson:"credits"`
	Trailers            []Trailer             `json:"trailers" bson:"trailers"`
}

type BelongsTo struct {
	ID           int64  `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	PosterPath   string `json:"poster_path" bson:"poster_path"`
	BackdropPath string `json:"backdrop_path" bson:"backdrop_path"`
}

type MovieGenres struct {
	ID   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type ProductionCompanies struct {
	ID            int64  `json:"id" bson:"id"`
	Name          string `json:"name" bson:"name"`
	LogoPath      string `json:"logo_path" bson:"logo_path"`
	OriginCountry string `json:"origin_country" bson:"origin_country"`
}

type ProductionCountries struct {
	ISO31661 string `json:"iso_3166_1" bson:"iso_3166_1"`
	Name     string `json:"name" bson:"name"`
}

type SpokenLanguages struct {
	EnglishName string `json:"english_name" bson:"english_name"`
	ISO6391     string `json:"iso_639_1" bson:"iso_639_1"`
	Name        string `json:"name" bson:"name"`
}

type MovieCredits struct {
	ID   int64       `json:"id" bson:"id"`
	Cast []MovieCrew `json:"cast" bson:"cast"`
	Crew []MovieCrew `json:"crew" bson:"crew"`
}

// lastest
type Trailer struct {
	ID          string    `json:"id" bson:"id"`
	ISO6391     string    `json:"iso_639_1" bson:"iso_639_1"`
	ISO31661    string    `json:"iso_3166_1" bson:"iso_3166_1"`
	Name        string    `json:"name" bson:"name"`
	Key         string    `json:"key" bson:"key"`
	Site        string    `json:"site" bson:"site"`
	Size        int64     `json:"size" bson:"size"`
	Type        string    `json:"type" bson:"type"`
	Official    bool      `json:"official" bson:"official"`
	PublishedAt time.Time `json:"published_at" bson:"published_at"`
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
