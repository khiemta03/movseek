package mongo

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"github.com/tmplam/movseek/pkg/mongo"
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
		return models.SavedItems{}, err
	}

	return s, err
}

func (repo implRepository) AddSavedItem(ctx context.Context, input saved_item.AddSavedItemOptions) error {
	col := repo.getSavedItemsCollection()

	model := repo.buildAddSavedItem(input)

	_, err := col.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (repo implRepository) UpdateSavedItem(ctx context.Context, input saved_item.UpdateSavedItemOptions) error {
	col := repo.getSavedItemsCollection()

	queryFilter := repo.buildUpdateSavedItemQuery(input)
	update := repo.buildUpdateSavedItem(input)

	_, err := col.UpdateOne(ctx, queryFilter, update)
	if err != nil {
		return err
	}

	return nil
}
