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

func (uc implUsecase) GetCommentsByMedia(ctx context.Context, input review.GetCommentsByMediaInput) (review.GetCommentsOutput, error) {
	comments, err := uc.repo.GetComments(ctx, review.Scope{
		MediaID: input.MediaID,
		Type:    input.Type,
	})
	if err != nil {
		return review.GetCommentsOutput{}, err
	}

	return review.GetCommentsOutput{
		MediaID:  input.MediaID,
		Type:     input.Type,
		Comments: comments,
	}, nil
}

func (uc implUsecase) GetCommentsByUserID(ctx context.Context, input review.GetCommentsByUserIDInput) (review.GetCommentsOutput, error) {
	comments, err := uc.repo.GetComments(ctx, review.Scope{
		UserID: input.UserID,
	})
	if err != nil {
		return review.GetCommentsOutput{}, err
	}

	return review.GetCommentsOutput{
		UserID:   input.UserID,
		Comments: comments,
	}, nil
}

func (uc implUsecase) GetRatingsByMedia(ctx context.Context, input review.GetRatingsByMediaInput) (review.GetRatingsOutput, error) {
	ratings, err := uc.repo.GetRatings(ctx, review.Scope{
		MediaID: input.MediaID,
		Type:    input.Type,
	})
	if err != nil {
		return review.GetRatingsOutput{}, err
	}

	return review.GetRatingsOutput{
		MediaID: input.MediaID,
		Type:    input.Type,
		Ratings: ratings,
	}, nil
}

func (uc implUsecase) GetRatingsByUserID(ctx context.Context, input review.GetRatingsByUserIDInput) (review.GetRatingsOutput, error) {
	ratings, err := uc.repo.GetRatings(ctx, review.Scope{
		UserID: input.UserID,
	})
	if err != nil {
		return review.GetRatingsOutput{}, err
	}

	return review.GetRatingsOutput{
		UserID:  input.UserID,
		Ratings: ratings,
	}, nil
}
