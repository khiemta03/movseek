package saved_item

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetSavedItem(ctx context.Context, input GetSavedItemFilter) (models.SavedItems, error)
	AddSavedItem(ctx context.Context, input AddSavedItemOptions) error
	UpdateSavedItem(ctx context.Context, input UpdateSavedItemOptions) error
}
