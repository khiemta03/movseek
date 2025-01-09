package http

import (
	"errors"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/review"
)

type addCommentRequest struct {
	UserID  string `form:"user_id" binding:"required"`
	MediaID int64  `form:"media_id" binding:"required"`
	Type    string `form:"type" binding:"required"`
	Comment string `form:"comment" binding:"required"`
}

func (req addCommentRequest) toInput() review.AddCommentInput {
	return review.AddCommentInput{
		Scope: review.Scope{
			UserID:  req.UserID,
			MediaID: req.MediaID,
			Type:    req.Type,
		},
		Comment: req.Comment,
	}
}

func (req addCommentRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}

	return nil
}

type addRatingRequest struct {
	UserID  string  `form:"user_id" binding:"required"`
	MediaID int64   `form:"media_id" binding:"required"`
	Type    string  `form:"type" binding:"required"`
	Rating  float64 `form:"rating" binding:"required"`
}

func (req addRatingRequest) toInput() review.AddRatingInput {
	return review.AddRatingInput{
		Scope: review.Scope{
			UserID:  req.UserID,
			MediaID: req.MediaID,
			Type:    req.Type,
		},
		Rating: req.Rating,
	}
}

func (req addRatingRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}

	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}

	if req.Rating < 0 || req.Rating > 10 {
		return errors.New("rating must be between 0 and 10")
	}

	return nil
}

type updateCommentRequest struct {
	UserID  string `uri:"user_id" binding:"required"`
	MediaID int64  `form:"media_id"`
	Type    string `form:"type"`
	Comment string `form:"comment"`
}

func (req updateCommentRequest) toInput() review.UpdateCommentInput {
	return review.UpdateCommentInput{
		Scope: review.Scope{
			UserID:  req.UserID,
			MediaID: req.MediaID,
			Type:    req.Type,
		},
		Comment: req.Comment,
	}
}

func (req updateCommentRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}

	return nil
}

type updateRatingRequest struct {
	UserID  string  `uri:"user_id" binding:"required"`
	MediaID int64   `form:"media_id"`
	Type    string  `form:"type"`
	Rating  float64 `form:"rating"`
}

func (req updateRatingRequest) toInput() review.UpdateRatingInput {
	return review.UpdateRatingInput{
		Scope: review.Scope{
			UserID:  req.UserID,
			MediaID: req.MediaID,
			Type:    req.Type,
		},
		Rating: req.Rating,
	}
}

func (req updateRatingRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}
	if req.Rating < 0 || req.Rating > 10 {
		return errors.New("rating must be between 0 and 10")
	}

	return nil
}

type deleteCommentRequest struct {
	UserID  string `uri:"user_id" binding:"required"`
	MediaID int64  `form:"media_id"`
	Type    string `form:"type"`
}

func (req deleteCommentRequest) toInput() review.Scope {
	return review.Scope{
		UserID:  req.UserID,
		MediaID: req.MediaID,
		Type:    req.Type,
	}
}

func (req deleteCommentRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}
	if req.MediaID == 0 {
		return errors.New("media_id is required")
	}
	return nil
}

type deleteRatingRequest struct {
	UserID  string `uri:"user_id" binding:"required"`
	MediaID int64  `form:"media_id"`
	Type    string `form:"type"`
}

func (req deleteRatingRequest) toInput() review.Scope {
	return review.Scope{
		UserID:  req.UserID,
		MediaID: req.MediaID,
		Type:    req.Type,
	}
}

func (req deleteRatingRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}
	if req.MediaID == 0 {
		return errors.New("media_id is required")
	}
	return nil
}

type getCommentsRequestByMediaID struct {
	MediaID int64  `uri:"media_id" binding:"required"`
	Type    string `form:"type"`
}

func (req getCommentsRequestByMediaID) toInput() review.GetCommentsByMediaInput {
	return review.GetCommentsByMediaInput{
		MediaID: req.MediaID,
		Type:    req.Type,
	}
}

func (req getCommentsRequestByMediaID) validate() error {
	if req.MediaID == 0 {
		return errors.New("media_id is required")
	}
	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}
	return nil
}

type getCommentsRequestByUserID struct {
	UserID string `uri:"user_id" binding:"required"`
}

func (req getCommentsRequestByUserID) toInput() review.GetCommentsByUserIDInput {
	return review.GetCommentsByUserIDInput{
		UserID: req.UserID,
	}
}

type getRatingsRequestByMediaID struct {
	MediaID int64  `uri:"media_id" binding:"required"`
	Type    string `form:"type"`
}

func (req getRatingsRequestByMediaID) toInput() review.GetRatingsByMediaInput {
	return review.GetRatingsByMediaInput{
		MediaID: req.MediaID,
		Type:    req.Type,
	}
}

func (req getRatingsRequestByMediaID) validate() error {
	if req.MediaID == 0 {
		return errors.New("media_id is required")
	}
	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}
	return nil
}

type getRatingsRequestByUserID struct {
	UserID string `uri:"user_id" binding:"required"`
}

func (req getRatingsRequestByUserID) toInput() review.GetRatingsByUserIDInput {
	return review.GetRatingsByUserIDInput{
		UserID: req.UserID,
	}
}
