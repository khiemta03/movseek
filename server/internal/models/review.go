package models

const (
	MovieType  = "movie"
	TVShowType = "tv_show"
)

type Favourite struct {
	UserID   string  `json:"user_id" bson:"user_id"`
	MovieID  []int64 `json:"movie_id" bson:"movie_id,omitempty"`
	TVShowID []int64 `json:"tv_show_id" bson:"tv_show_id,omitempty"`
}

type Watchlist struct {
	UserID   string  `json:"user_id" bson:"user_id"`
	MovieID  []int64 `json:"movie_id" bson:"movie_id,omitempty"`
	TVShowID []int64 `json:"tv_show_id" bson:"tv_show_id,omitempty"`
}

type Review struct {
	UserID  string   `json:"user_id" bson:"user_id,omitempty"`
	MediaID int64    `json:"media_id" bson:"media_id,omitempty"`
	Type    string   `json:"type" bson:"type,omitempty"`
	Rating  *float64 `json:"rating" bson:"rating,omitempty"` // 0-10
	Comment *string  `json:"comment" bson:"comment,omitempty"`
}

// /reviews/tv
// /

// /reviews/media/:media_id?type=movie
// /reviews/media/:media_id?type=tv_show
// /reviews/person/:user_id
