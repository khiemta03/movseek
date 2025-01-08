package mongo

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie/repository"
	"github.com/tmplam/movseek/internal/person"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"github.com/tmplam/movseek/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	savedItemsCollection = "saved_items"
)

func (repo implRepository) getSavedItemsCollection() mongo.Collection {
	return repo.db.Collection(savedItemsCollection)
}

func (repo implRepository) GetSavedItem(ctx context.Context, input saved_item.GetSavedItemFilter) (models.SavedItems, error) {
	col := repo.getSavedItemsCollection()

	queryFilter := repo.buildSavedItemQuery(input)

	var s = models.SavedItems{}
	err := col.FindOne(ctx, queryFilter).Decode(&s)
	if err != nil {
		return models.SavedItems{}, repository.MapError(err)
	}

	return s, err
}

func (repo implRepository) AddToSavedItem(ctx context.Context, input saved_item.AddToSavedItemInput) error {
	col := repo.getSavedItemsCollection()

	queryFilter := repo.buildSavedItemQuery(input)

	findOptions := repo.buildGetPeopleOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	var people []models.PersonSummary
	err = cursor.All(ctx, &people)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	return people, nil
}

func (repo implRepository) GetPopularPeople(ctx context.Context, input person.GetPopularPeopleOptions) ([]models.PersonSummary, error) {
	col := repo.getPopularCollection()

	findOptions := repo.buildGetPeopleOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	var people []models.PersonSummary
	err = cursor.All(ctx, &people)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	return people, nil
}
