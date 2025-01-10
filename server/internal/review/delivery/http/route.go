package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	comments := r.Group("/comments")
	comments.GET("/media/:media_id", h.getCommentsByMedia)
	comments.GET("/user/:user_id", h.getCommentsByUser)
	comments.POST("/", h.addComment)
	comments.PUT("/user/:user_id", h.updateComment)
	comments.DELETE("/user/:user_id", h.deleteComment)

	ratings := r.Group("/ratings")
	ratings.GET("/media/:media_id", h.getRatingsByMedia)
	ratings.GET("/user/:user_id", h.getRatingsByUser)
	ratings.POST("/", h.addRating)
	ratings.PUT("/user/:user_id", h.updateRating)
	ratings.DELETE("/user/:user_id", h.deleteRating)
}
