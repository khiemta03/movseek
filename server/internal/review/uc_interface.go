package review

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Usecase interface {
	AddComment(ctx context.Context, input AddCommentInput) (models.Comment, error)
	UpdateComment(ctx context.Context, input UpdateCommentInput) error
	DeleteComment(ctx context.Context, input Scope) error
	AddRating(ctx context.Context, input AddRatingInput) (models.Rating, error)
	UpdateRating(ctx context.Context, input UpdateRatingInput) error
	DeleteRating(ctx context.Context, input Scope) error
	GetCommentsByMedia(ctx context.Context, input GetCommentsByMediaInput) (GetCommentsOutput, error)
	GetCommentsByUserID(ctx context.Context, input GetCommentsByUserIDInput) (GetCommentsOutput, error)
	GetRatingsByMedia(ctx context.Context, input GetRatingsByMediaInput) (GetRatingsOutput, error)
	GetRatingsByUserID(ctx context.Context, input GetRatingsByUserIDInput) (GetRatingsOutput, error)
}
