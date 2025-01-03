package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo implRepository) buildMovieQuery(movieID string) (bson.M, error) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return bson.M{}, err
	}

	queryFilter := bson.M{
		"_id": id,
	}

	return queryFilter, nil
}
