package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	r.GET("/search/person", h.searchPeople)

	person := r.Group("/person")
	person.GET("/:id", h.getOnePerson)
	person.GET("/trending/:type", h.getTrendingPeople)
	person.GET("/popular", h.getPopularPeople)
	person.GET("/:id/credits/movie", h.getPersonCredits)
	person.GET("/:id/credits/tv", h.getTVCredits)
}
