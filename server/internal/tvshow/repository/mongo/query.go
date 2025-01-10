package mongo

import (
	"github.com/tmplam/movseek/internal/tvshow"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo implRepository) buildTVShowQuery(tvID int64) (bson.M, error) {
	queryFilter := bson.M{
		"id": tvID,
	}

	return queryFilter, nil
}

func (repo implRepository) buildListTVShowsQuery(input tvshow.ListTVsOptions) bson.M {
	queryFilter := bson.M{}
	if input.Query != "" {
		queryFilter["title"] = bson.M{
			"$regex":   input.Query,
			"$options": "i",
		}
	}

	return queryFilter
}

func (repo implRepository) buildGetTVShowFindOptions(input tvshow.GetTVFilter) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))
	findOptions.SetSort(bson.M{"release_date": -1})
	return findOptions
}
