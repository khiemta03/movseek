package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
	"github.com/tmplam/movseek/internal/saved-item/repository"
)

func (uc implUsecase) GetSavedItem(ctx context.Context, input saved_item.GetSavedItemFilter) (models.SavedItems, error) {
	return uc.repo.GetSavedItem(ctx, input)
}

func (uc implUsecase) AddToSavedItem(ctx context.Context, input saved_item.AddToSavedItemInput) error {
	_, err := uc.repo.GetSavedItem(ctx, saved_item.GetSavedItemFilter{
		UserID: input.UserID,
		Type:   input.SaveItemType,
	})
	if err != nil {
		if err == repository.ErrNotFound {
			err := uc.repo.AddSavedItem(ctx, saved_item.AddSavedItemOptions{
				UserID:       input.UserID,
				MediaID:      input.MediaID,
				SaveItemType: input.SaveItemType,
				MediaType:    input.MediaType,
			})
			if err != nil {
				return err
			}
		}
		return err
	}

	options := saved_item.UpdateSavedItemOptions{
		UserID:       input.UserID,
		MediaID:      input.MediaID,
		SaveItemType: input.SaveItemType,
		MediaType:    input.MediaType,
		Action:       "add",
	}
	return uc.repo.UpdateSavedItem(ctx, options)
}

func (uc implUsecase) RemoveFromSavedItem(ctx context.Context, input saved_item.RemoveFromSavedItemInput) error {
	options := saved_item.UpdateSavedItemOptions{
		UserID:       input.UserID,
		MediaID:      input.MediaID,
		SaveItemType: input.Type,
		MediaType:    input.MediaType,
		Action:       "remove",
	}
	return uc.repo.UpdateSavedItem(ctx, options)
}
