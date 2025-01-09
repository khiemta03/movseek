package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"github.com/tmplam/movseek/pkg/mongo"
)

func (uc implUsecase) GetSavedItem(ctx context.Context, input saved_item.GetSavedItemFilter) (models.SavedItems, error) {
	return uc.repo.GetSavedItem(ctx, input)
}

func (uc implUsecase) AddToSavedItem(ctx context.Context, input saved_item.AddToSavedItemInput) error {
	if input.MediaType == models.MovieType {
		_, err := uc.movieUC.GetOneMovie(ctx, input.MediaID)
		if err != nil {
			return err
		}
	} else if input.MediaType == models.TVShowType {
		_, err := uc.tvUC.GetOneTVShow(ctx, input.MediaID)
		if err != nil {
			return err
		}
	}

	_, err := uc.repo.GetSavedItem(ctx, saved_item.GetSavedItemFilter{
		UserID: input.UserID,
		Type:   input.SavedItemType,
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err := uc.repo.AddSavedItem(ctx, saved_item.AddSavedItemOptions{
				UserID:        input.UserID,
				MediaID:       input.MediaID,
				SavedItemType: input.SavedItemType,
				MediaType:     input.MediaType,
			})
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	options := saved_item.UpdateSavedItemOptions{
		UserID:        input.UserID,
		MediaID:       input.MediaID,
		SavedItemType: input.SavedItemType,
		MediaType:     input.MediaType,
		Action:        "add",
	}
	return uc.repo.UpdateSavedItem(ctx, options)
}

func (uc implUsecase) RemoveFromSavedItem(ctx context.Context, input saved_item.RemoveFromSavedItemInput) error {
	options := saved_item.UpdateSavedItemOptions{
		UserID:        input.UserID,
		MediaID:       input.MediaID,
		SavedItemType: input.Type,
		MediaType:     input.MediaType,
		Action:        "remove",
	}
	return uc.repo.UpdateSavedItem(ctx, options)
}
