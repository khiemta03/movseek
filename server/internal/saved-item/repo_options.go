package saved_item

type ListSavedItemsOptions struct {
	Query  string
	Filter GetSavedItemFilter
}

type AddSavedItemOptions struct {
	UserID       int64  `json:"user_id"`
	MediaID      int64  `json:"media_id"`
	SaveItemType string `json:"type"`
	MediaType    string `json:"media_type"`
}

type UpdateSavedItemOptions struct {
	UserID       int64  `json:"user_id"`
	MediaID      int64  `json:"media_id"`
	SaveItemType string `json:"type"`
	MediaType    string `json:"media_type"`
	Action       string `json:"action"`
}
