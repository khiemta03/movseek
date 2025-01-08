package models

const (
	SaveItemTypeFavorite  = "favorite"
	SaveItemTypeWatchlist = "watchlist"
)

type SavedItems struct {
	UserID   int64   `json:"user_id" bson:"user_id"`
	MovieID  []int64 `json:"movie_id" bson:"movie_id,omitempty"`
	TVShowID []int64 `json:"tv_show_id" bson:"tv_show_id,omitempty"`
	Type     string  `json:"type" bson:"type"`
}
