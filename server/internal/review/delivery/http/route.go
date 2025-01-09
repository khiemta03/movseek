package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	comments := r.Group("/comments")
	comments.GET("/:media_id", h.getCommentsByMedia)
	comments.GET("/:user_id", h.getCommentsByUser)
	comments.POST("/", h.addComment)
	comments.PUT("/:user_id", h.updateComment)
	comments.DELETE("/:user_id", h.deleteComment)

	ratings := r.Group("/ratings")
	ratings.GET("/:media_id", h.getRatingsByMedia)
	ratings.GET("/:user_id", h.getRatingsByUser)
	ratings.POST("/", h.addRating)
	ratings.PUT("/:user_id", h.updateRating)
	ratings.DELETE("/:user_id", h.deleteRating)
}
