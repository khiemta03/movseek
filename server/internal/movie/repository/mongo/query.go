package mongo

import (
	"github.com/tmplam/movseek/internal/movie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo implRepository) buildMovieQuery(movieID int64) (bson.M, error) {
	queryFilter := bson.M{
		"id": movieID,
	}

	return queryFilter, nil
}

func (repo implRepository) buildListMoviesQuery(input movie.ListMoviesOptions) bson.M {
	queryFilter := bson.M{
		"title": bson.M{
			"$regex":   input.Query,
			"$options": "i",
		},
	}

	return queryFilter
}

func (repo implRepository) buildGetMovieFindOptions(input movie.GetMovieFilter) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))
	findOptions.SetSort(bson.M{"release_date": -1})
	return findOptions
}
