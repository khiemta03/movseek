package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TVGenre struct {
	ObjectID primitive.ObjectID `json:"_id" bson:"_id"`
	TMDBID   int                `json:"tmdb_id" bson:"tmdb_id"`
	ID       int                `json:"id" bson:"id"`
	Name     string             `json:"name" bson:"name"`
}

