package saved_item

type GetSavedItemFilter struct {
	UserID  int64  `json:"user_id"`
	Type    string `json:"type"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

type AddToSavedItemInput struct {
	UserID       int64  `json:"user_id"`
	MediaID      int64  `json:"media_id"`
	SaveItemType string `json:"type"`
	MediaType    string `json:"media_type"`
}

type RemoveFromSavedItemInput struct {
	UserID    int64  `json:"user_id"`
	MediaID   int64  `json:"media_id"`
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
}
