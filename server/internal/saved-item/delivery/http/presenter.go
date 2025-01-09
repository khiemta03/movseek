package http

import (
	"errors"

	"github.com/tmplam/movseek/internal/models"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
)

type getSavedItemRequest struct {
	UserID string `uri:"user_id" binding:"required"`
	Type   string `form:"type"`
}

func (req getSavedItemRequest) toInput() saved_item.GetSavedItemFilter {
	return saved_item.GetSavedItemFilter{
		UserID: req.UserID,
		Type:   req.Type,
	}
}

func (req getSavedItemRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}

	if req.Type != models.SavedItemTypeFavorite && req.Type != models.SavedItemTypeWatchlist {
		return errors.New("type must be favorite or watchlist")
	}

	return nil
}

type addToSavedItemRequest struct {
	UserID        string `json:"user_id" binding:"required"`
	SavedItemType string `json:"saved_item_type" binding:"required"`
	MediaID       int64  `json:"media_id" binding:"required"`
	MediaType     string `json:"media_type" binding:"required"`
}

func (req addToSavedItemRequest) toInput() saved_item.AddToSavedItemInput {
	return saved_item.AddToSavedItemInput{
		UserID:        req.UserID,
		SavedItemType: req.SavedItemType,
		MediaID:       req.MediaID,
		MediaType:     req.MediaType,
	}
}

func (req addToSavedItemRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.SavedItemType != models.SavedItemTypeFavorite && req.SavedItemType != models.SavedItemTypeWatchlist {
		return errors.New("saved_item_type must be favorite or watchlist")
	}

	if req.MediaType != models.MovieType && req.MediaType != models.TVShowType {
		return errors.New("media_type must be movie or tv_show")
	}

	return nil
}

type removeFromSavedItemRequest struct {
	UserID        string `json:"user_id" binding:"required"`
	SavedItemType string `json:"saved_item_type" binding:"required"`
	MediaID       int64  `json:"media_id" binding:"required"`
	MediaType     string `json:"media_type" binding:"required"`
}

func (req removeFromSavedItemRequest) toInput() saved_item.RemoveFromSavedItemInput {
	return saved_item.RemoveFromSavedItemInput{
		UserID:    req.UserID,
		Type:      req.SavedItemType,
		MediaID:   req.MediaID,
		MediaType: req.MediaType,
	}
}

func (req removeFromSavedItemRequest) validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.SavedItemType != models.SavedItemTypeFavorite && req.SavedItemType != models.SavedItemTypeWatchlist {
		return errors.New("saved_item_type must be favorite or watchlist")
	}
	if req.MediaType != models.MovieType && req.MediaType != models.TVShowType {
		return errors.New("media_type must be movie or tv_show")
	}
	return nil
}
