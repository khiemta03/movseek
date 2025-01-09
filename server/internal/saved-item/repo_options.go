package saved_item

type GetSavedItemsOptions struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

type AddSavedItemOptions struct {
	UserID        string `json:"user_id"`
	MediaID       int64  `json:"media_id"`
	SavedItemType string `json:"type"`
	MediaType     string `json:"media_type"`
}

type UpdateSavedItemOptions struct {
	UserID        string `json:"user_id"`
	MediaID       int64  `json:"media_id"`
	SavedItemType string `json:"type"`
	MediaType     string `json:"media_type"`
	Action        string `json:"action"`
}
