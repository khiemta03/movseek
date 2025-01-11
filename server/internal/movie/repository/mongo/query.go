package mongo

import (
	"github.com/tmplam/movseek/internal/movie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo implRepository) buildMovieQuery(movieID int64) (bson.M, error) {
	queryFilter := bson.M{
		"id": movieID,
	}

	return queryFilter, nil
}

func (repo implRepository) buildListMoviesQuery(input movie.ListMoviesOptions) (bson.M, error) {
	queryFilter := bson.M{}
	if input.Query != "" {
		queryFilter["$or"] = []bson.M{
			{
				"title": bson.M{
					"$regex":   input.Query,
					"$options": "i",
				},
			},
			{
				"keywords.name": bson.M{
					"$regex":   input.Query,
					"$options": "i",
				},
			},
			{
				"overview": bson.M{
					"$regex":   input.Query,
					"$options": "i",
				},
			},
			{
				"belongs_to_collection.name": bson.M{
					"$regex":   input.Query,
					"$options": "i",
				},
			},
		}
	}

	if len(input.Filter.IDs) > 0 {
		queryFilter["id"] = bson.M{"$in": input.Filter.IDs}
	}

	if len(input.Filter.ObjectIDs) > 0 {
		objectIDs := make([]primitive.ObjectID, len(input.Filter.ObjectIDs))
		for i, id := range input.Filter.ObjectIDs {
			objectID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return bson.M{}, err
			}
			objectIDs[i] = objectID
		}
		queryFilter["_id"] = bson.M{"$in": objectIDs}
	}

	return queryFilter, nil
}

func (repo implRepository) buildGetMovieFindOptions(input movie.GetMovieFilter) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))
	findOptions.SetSort(bson.M{"release_date": -1})
	return findOptions
}
