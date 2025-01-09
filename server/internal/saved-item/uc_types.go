package saved_item

type GetSavedItemFilter struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

type AddToSavedItemInput struct {
	UserID        string `json:"user_id"`
	MediaID       int64  `json:"media_id"`
	SavedItemType string `json:"type"`
	MediaType     string `json:"media_type"`
}

type RemoveFromSavedItemInput struct {
	UserID    string `json:"user_id"`
	MediaID   int64  `json:"media_id"`
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
}
