package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	ai := r.Group("/ai")
	ai.GET("/health", h.healthy)
	ai.GET("/retriever", h.retriever)
	ai.POST("/navigate", h.navigate)
}
