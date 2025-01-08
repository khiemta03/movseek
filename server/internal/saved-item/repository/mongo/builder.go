package saved_item

import (
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo implRepository) buildSavedItemModel(input saved_item.AddToSavedItemInput) bson.M {
	queryFilter := bson.M{
		"user_id": input.UserID,
		"type":    input.Type,
	}

	return queryFilter
}
