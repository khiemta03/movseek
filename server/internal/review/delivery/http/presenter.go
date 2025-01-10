package http

import (
	"errors"
	"time"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/review"
)

type addCommentRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	MediaID int64  `json:"media_id" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Comment string `json:"comment" binding:"required"`
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
	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
	}

	return nil
}

type addRatingRequest struct {
	UserID  string  `json:"user_id" binding:"required"`
	MediaID int64   `json:"media_id" binding:"required"`
	Type    string  `json:"type" binding:"required"`
	Rating  float64 `json:"rating" binding:"required"`
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

	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
	}

	if req.Rating < 0 || req.Rating > 10 {
		return errors.New("rating must be between 0 and 10")
	}

	return nil
}

type updateCommentRequest struct {
	UserID  string `uri:"user_id" binding:"required"`
	MediaID int64  `json:"media_id"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
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
	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
	}

	return nil
}

type updateRatingRequest struct {
	UserID  string  `uri:"user_id" binding:"required"`
	MediaID int64   `json:"media_id"`
	Type    string  `json:"type"`
	Rating  float64 `json:"rating"`
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
	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
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
	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
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
	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
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
	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
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
	if req.Type != models.MovieType && req.Type != models.TVShowType {
		return errors.New("type must be movie or tv")
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

type getCommentsByMediaDetail struct {
	UserID    string     `json:"user_id"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type getCommentsByMediaDetailsResponse struct {
	MediaID  int64                      `json:"media_id"`
	Type     string                     `json:"type"`
	Comments []getCommentsByMediaDetail `json:"comments"`
}

func (h handlerImpl) getCommentsByMediaResponse(output review.GetCommentsOutput) getCommentsByMediaDetailsResponse {
	res := getCommentsByMediaDetailsResponse{
		MediaID: output.MediaID,
		Type:    output.Type,
	}

	res.Comments = make([]getCommentsByMediaDetail, 0, len(output.Comments))
	for _, comment := range output.Comments {
		res.Comments = append(res.Comments, getCommentsByMediaDetail{
			UserID:    comment.UserID,
			Comment:   comment.Comment,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	return res
}

type getRatingsByMediaDetail struct {
	UserID    string     `json:"user_id"`
	Rating    float64    `json:"rating"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type getRatingsByMediaDetailsResponse struct {
	MediaID int64                     `json:"media_id"`
	Type    string                    `json:"type"`
	Average float64                   `json:"average"`
	Count   int                       `json:"count"`
	Ratings []getRatingsByMediaDetail `json:"ratings"`
}

func (h handlerImpl) getRatingsByMediaResponse(output review.GetRatingsOutput) getRatingsByMediaDetailsResponse {
	res := getRatingsByMediaDetailsResponse{
		MediaID: output.MediaID,
		Type:    output.Type,
		Count:   len(output.Ratings),
	}

	total := 0.0
	res.Ratings = make([]getRatingsByMediaDetail, 0, len(output.Ratings))
	for _, rating := range output.Ratings {
		res.Ratings = append(res.Ratings, getRatingsByMediaDetail{
			UserID:    rating.UserID,
			Rating:    rating.Rating,
			CreatedAt: rating.CreatedAt,
			UpdatedAt: rating.UpdatedAt,
		})
		total += rating.Rating
	}

	if len(output.Ratings) > 0 {
		res.Average = total / float64(len(output.Ratings))
	}

	return res
}

type getCommentsByUserDetail struct {
	MediaID   int64      `json:"media_id"`
	Type      string     `json:"type"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type getCommentsByUserDetailsResponse struct {
	UserID   string                    `json:"user_id"`
	Comments []getCommentsByUserDetail `json:"comments"`
}

func (h handlerImpl) getCommentsByUserResponse(output review.GetCommentsOutput) getCommentsByUserDetailsResponse {
	res := getCommentsByUserDetailsResponse{
		UserID: output.UserID,
	}

	res.Comments = make([]getCommentsByUserDetail, 0, len(output.Comments))
	for _, comment := range output.Comments {
		res.Comments = append(res.Comments, getCommentsByUserDetail{
			MediaID:   comment.MediaID,
			Type:      comment.Type,
			Comment:   comment.Comment,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	return res
}

type getRatingsByUserDetail struct {
	MediaID   int64      `json:"media_id"`
	Type      string     `json:"type"`
	Rating    float64    `json:"rating"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type getRatingsByUserDetailsResponse struct {
	UserID  string                   `json:"user_id"`
	Count   int                      `json:"count"`
	Ratings []getRatingsByUserDetail `json:"ratings"`
}

func (h handlerImpl) getRatingsByUserResponse(output review.GetRatingsOutput) getRatingsByUserDetailsResponse {
	res := getRatingsByUserDetailsResponse{
		UserID: output.UserID,
		Count:  len(output.Ratings),
	}

	res.Ratings = make([]getRatingsByUserDetail, 0, len(output.Ratings))
	for _, rating := range output.Ratings {
		res.Ratings = append(res.Ratings, getRatingsByUserDetail{
			MediaID:   rating.MediaID,
			Type:      rating.Type,
			Rating:    rating.Rating,
			CreatedAt: rating.CreatedAt,
			UpdatedAt: rating.UpdatedAt,
		})
	}

	return res
}
