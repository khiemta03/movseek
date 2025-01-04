package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	r.GET("/search/movie", h.searchMovies)

	movie := r.Group("/movie")
	movie.GET("/:id", h.getOneMovie)
	movie.GET("/:id/credits", h.getMovieCredits)
}
