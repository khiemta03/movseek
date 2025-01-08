package mongo

import (
	"github.com/tmplam/movseek/internal/person"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo implRepository) buildSavedItemQuery(input saved_item.GetSavedItemFilter) bson.M {
	queryFilter := bson.M{
		"user_id": input.UserID,
		"type":    input.Type,
	}

	return queryFilter
}

func (rep) 

func (repo implRepository) buildGetPeopleOptions(input person.GetPersonFilter) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))

	return findOptions
}
