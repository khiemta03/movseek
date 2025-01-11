package review

import "github.com/tmplam/movseek/internal/models"

type Scope struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"`
	Avatar   string `json:"avatar"`
	MediaID  int64  `json:"media_id"`
	Type     string `json:"type"`
}

type AddCommentInput struct {
	Scope
	Comment string `json:"comment"`
}

type AddRatingInput struct {
	Scope
	Rating float64 `json:"rating"`
}

type UpdateCommentInput struct {
	Scope
	Comment string `json:"comment"`
}

type UpdateRatingInput struct {
	Scope
	Rating float64 `json:"rating"`
}

type GetCommentsByMediaInput struct {
	MediaID int64  `json:"media_id"`
	Type    string `json:"type"`
}

type GetCommentsByUserIDInput struct {
	UserID string `json:"user_id"`
}

type GetRatingsByMediaInput struct {
	MediaID int64  `json:"media_id"`
	Type    string `json:"type"`
}

type GetRatingsByUserIDInput struct {
	UserID string `json:"user_id"`
}

type GetCommentsOutput struct {
	UserID   string           `json:"user_id"`
	UserName string           `json:"username"`
	Avatar   string           `json:"avatar"`
	MediaID  int64            `json:"media_id"`
	Type     string           `json:"type"`
	Comments []models.Comment `json:"comments"`
}

type GetRatingsOutput struct {
	UserID   string          `json:"user_id"`
	UserName string          `json:"username"`
	Avatar   string          `json:"avatar"`
	MediaID  int64           `json:"media_id"`
	Type     string          `json:"type"`
	Ratings  []models.Rating `json:"ratings"`
}
