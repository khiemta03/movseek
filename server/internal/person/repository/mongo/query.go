package mongo

import (
	"github.com/tmplam/movseek/internal/person"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo implRepository) buildPersonQuery(personID string) (bson.M, error) {
	id, err := primitive.ObjectIDFromHex(personID)
	if err != nil {
		return bson.M{}, err
	}

	queryFilter := bson.M{
		"_id": id,
	}

	return queryFilter, nil
}

func (repo implRepository) buildListPeopleQuery(input person.ListPeopleOptions) bson.M {
	queryFilter := bson.M{
		"name": bson.M{
			"$regex":   input.Query,
			"$options": "i",
		},
	}

	return queryFilter
}

func (repo implRepository) buildGetPeopleOptions(input person.GetPersonFilter) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))

	return findOptions
}
