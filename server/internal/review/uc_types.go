package review

type Scope struct {
	UserID  string `json:"user_id"`
	MediaID int64  `json:"media_id"`
	Type    string `json:"type"`
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
