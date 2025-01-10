package mongo

import (
	"github.com/tmplam/movseek/internal/models"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo implRepository) buildSavedItemQuery(input saved_item.GetSavedItemFilter) bson.M {
	queryFilter := bson.M{
		"user_id": input.UserID,
		"type":    input.Type,
	}

	return queryFilter
}

func (repo implRepository) buildUpdateSavedItemQuery(input saved_item.UpdateSavedItemOptions) bson.M {
	queryFilter := bson.M{
		"user_id": input.UserID,
		"type":    input.SavedItemType,
	}

	if input.Action == "add" {
		field := "movie_id"
		if input.MediaType == models.TVShowType {
			field = "tv_show_id"
		}

		queryFilter[field] = bson.M{
			"$nin": []int64{input.MediaID},
		}
	}

	return queryFilter
}
