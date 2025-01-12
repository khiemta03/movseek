package http

type retrieverRequest struct {
	LLMAPIKey      string  `form:"llm_api_key" binding:"required"`
	Query          string  `form:"query" binding:"required"`
	CollectionName string  `form:"collection_name" binding:"required"`
	Amount         int     `form:"amount"`
	Threshold      float64 `form:"threshold"`
}

type navigateRequest struct {
	LLMAPIKey string `form:"llm_api_key" binding:"required"`
	Query     string `form:"query" binding:"required"`
}
