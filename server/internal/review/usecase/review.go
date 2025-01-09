package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/review"
)

func (uc implUsecase) AddComment(ctx context.Context, input review.AddCommentInput) (models.Comment, error) {
	return uc.repo.AddComment(ctx, input)
}

func (uc implUsecase) UpdateComment(ctx context.Context, input review.UpdateCommentInput) error {
	return uc.repo.UpdateComment(ctx, input)
}

func (uc implUsecase) DeleteComment(ctx context.Context, input review.Scope) error {
	return uc.repo.DeleteComment(ctx, input)
}

func (uc implUsecase) AddRating(ctx context.Context, input review.AddRatingInput) (models.Rating, error) {
	return uc.repo.AddRating(ctx, input)
}

func (uc implUsecase) UpdateRating(ctx context.Context, input review.UpdateRatingInput) error {
	return uc.repo.UpdateRating(ctx, input)
}

func (uc implUsecase) DeleteRating(ctx context.Context, input review.Scope) error {
	return uc.repo.DeleteRating(ctx, input)
}

func (uc implUsecase) GetCommentsByMedia(ctx context.Context, input review.GetCommentsByMediaInput) ([]models.Comment, error) {
	return uc.repo.GetComments(ctx, review.Scope{
		MediaID: input.MediaID,
		Type:    input.Type,
	})
}

func (uc implUsecase) GetCommentsByUserID(ctx context.Context, input review.GetCommentsByUserIDInput) ([]models.Comment, error) {
	return uc.repo.GetComments(ctx, review.Scope{
		UserID: input.UserID,
	})
}

func (uc implUsecase) GetRatingsByMedia(ctx context.Context, input review.GetRatingsByMediaInput) ([]models.Rating, error) {
	return uc.repo.GetRatings(ctx, review.Scope{
		MediaID: input.MediaID,
		Type:    input.Type,
	})
}

func (uc implUsecase) GetRatingsByUserID(ctx context.Context, input review.GetRatingsByUserIDInput) ([]models.Rating, error) {
	return uc.repo.GetRatings(ctx, review.Scope{
		UserID: input.UserID,
	})
}
