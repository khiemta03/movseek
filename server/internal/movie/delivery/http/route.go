package http

import "github.com/gin-gonic/gin"

func (h *handler) MapRoutes(r *gin.RouterGroup) {
	r.GET("/:id", h.getOneMovie)
	r.GET("/:id/credits", h.getMovieCredits)
}
