package review

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	AddComment(ctx context.Context, input AddCommentInput) (models.Comment, error)
	UpdateComment(ctx context.Context, input UpdateCommentInput) error
	DeleteComment(ctx context.Context, input Scope) error
	AddRating(ctx context.Context, input AddRatingInput) (models.Rating, error)
	UpdateRating(ctx context.Context, input UpdateRatingInput) error
	DeleteRating(ctx context.Context, input Scope) error
	GetComments(ctx context.Context, input Scope) ([]models.Comment, error)
	GetRatings(ctx context.Context, input Scope) ([]models.Rating, error)
}
