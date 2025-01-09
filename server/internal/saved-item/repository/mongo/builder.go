package mongo

import (
	"github.com/tmplam/movseek/internal/models"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo implRepository) buildUpdateSavedItem(input saved_item.UpdateSavedItemOptions) bson.M {
	update := bson.M{}

	if input.Action == "add" {
		pushField := bson.M{}
		if input.MediaType == models.MovieType {
			pushField["movie_id"] = input.MediaID
		} else if input.MediaType == models.TVShowType {
			pushField["tv_show_id"] = input.MediaID
		}
		update["$push"] = pushField
	} else if input.Action == "remove" {
		pullField := bson.M{}
		if input.MediaType == models.MovieType {
			pullField["movie_id"] = input.MediaID
		} else if input.MediaType == models.TVShowType {
			pullField["tv_show_id"] = input.MediaID
		}
		update["$pull"] = pullField
	}

	return update
}

func (repo implRepository) buildAddSavedItem(input saved_item.AddSavedItemOptions) models.SavedItems {
	var savedItem models.SavedItems

	savedItem.UserID = input.UserID
	savedItem.Type = input.SaveItemType

	if input.MediaType == models.MovieType {
		savedItem.MovieID = append(savedItem.MovieID, input.MediaID)
	} else if input.MediaType == models.TVShowType {
		savedItem.TVShowID = append(savedItem.TVShowID, input.MediaID)
	}

	return savedItem
}
