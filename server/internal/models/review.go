package models

import "time"

const (
	MovieType  = "movie"
	TVShowType = "tv_show"
)

type Comment struct {
	UserID    string     `json:"user_id" bson:"user_id,omitempty"`
	UserName  string     `json:"username" bson:"username,omitempty"`
	Avatar    string     `json:"avatar" bson:"avatar,omitempty"`
	MediaID   int64      `json:"media_id" bson:"media_id,omitempty"`
	Type      string     `json:"type" bson:"type,omitempty"`
	Comment   string     `json:"comment" bson:"comment,omitempty"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Rating struct {
	UserID    string     `json:"user_id" bson:"user_id,omitempty"`
	UserName  string     `json:"username" bson:"username,omitempty"`
	Avatar    string     `json:"avatar" bson:"avatar,omitempty"`
	MediaID   int64      `json:"media_id" bson:"media_id,omitempty"`
	Type      string     `json:"type" bson:"type,omitempty"`
	Rating    float64    `json:"rating" bson:"rating,omitempty"` // 0-10
	CreatedAt time.Time  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
