package models

const (
	MovieType  = "movie"
	TVShowType = "tv_show"
)

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
