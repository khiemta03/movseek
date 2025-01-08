package saved_item

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetSavedItem(ctx context.Context, input GetSavedItemFilter) (models.SavedItems, error)
	AddToSavedItem(ctx context.Context, input AddToSavedItemInput) error
	RemoveFromSavedItem(ctx context.Context, input RemoveFromSavedItemInput) error
}
