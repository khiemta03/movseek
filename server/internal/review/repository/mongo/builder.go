package mongo

import (
	"time"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/review"
)

func (repo implRepository) buildAddComment(input review.AddCommentInput) models.Comment {
	comment := models.Comment{
		UserID:    input.UserID,
		UserName:  input.UserName,
		Avatar:    input.Avatar,
		MediaID:   input.MediaID,
		Type:      input.Type,
		Comment:   input.Comment,
		CreatedAt: time.Now(),
	}

	return comment
}

func (repo implRepository) buildAddRating(input review.AddRatingInput) models.Rating {
	rating := models.Rating{
		UserID:    input.UserID,
		UserName:  input.UserName,
		Avatar:    input.Avatar,
		MediaID:   input.MediaID,
		Type:      input.Type,
		Rating:    input.Rating,
		CreatedAt: time.Now(),
	}

	return rating
}

func (repo implRepository) buildUpdateComment(input review.UpdateCommentInput) models.Comment {
	now := time.Now()
	comment := models.Comment{
		UserID:    input.UserID,
		UserName:  input.UserName,
		Avatar:    input.Avatar,
		MediaID:   input.MediaID,
		Type:      input.Type,
		Comment:   input.Comment,
		UpdatedAt: &now,
	}

	return comment
}

func (repo implRepository) buildUpdateRating(input review.UpdateRatingInput) models.Rating {
	now := time.Now()
	rating := models.Rating{
		UserID:    input.UserID,
		UserName:  input.UserName,
		Avatar:    input.Avatar,
		MediaID:   input.MediaID,
		Type:      input.Type,
		Rating:    input.Rating,
		UpdatedAt: &now,
	}

	return rating
}
