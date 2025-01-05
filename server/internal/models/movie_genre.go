package models

type MovieGenre struct {
	ID      int    `json:"_id"`
	TMDBID  int    `json:"tmdb_id"`
	MovieID int    `json:"id"`
	Name    string `json:"name"`
}
