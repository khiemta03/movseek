package mongo

import (
	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/review"
)

func (repo implRepository) buildAddComment(input review.AddCommentInput) models.Comment {
	comment := models.Comment{
		UserID:  input.UserID,
		MediaID: input.MediaID,
		Type:    input.Type,
		Comment: input.Comment,
	}

	return comment
}

func (repo implRepository) buildAddRating(input review.AddRatingInput) models.Rating {
	rating := models.Rating{
		UserID:  input.UserID,
		MediaID: input.MediaID,
		Type:    input.Type,
		Rating:  input.Rating,
	}

	return rating
}

func (repo implRepository) buildUpdateComment(input review.UpdateCommentInput) models.Comment {
	comment := models.Comment{
		UserID:  input.UserID,
		MediaID: input.MediaID,
		Type:    input.Type,
		Comment: input.Comment,
	}

	return comment
}

func (repo implRepository) buildUpdateRating(input review.UpdateRatingInput) models.Rating {
	rating := models.Rating{
		UserID:  input.UserID,
		MediaID: input.MediaID,
		Type:    input.Type,
		Rating:  input.Rating,
	}

	return rating
}
