package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	r.GET("/search/tv", h.searchTVShows)

	tvshow := r.Group("/tvshow")
	tvshow.GET("/:id", h.getOneTVShow)
	tvshow.GET("/upcoming", h.getUpcomingTVShows)
	tvshow.GET("/trending/:type", h.getTrendingTVShows)
	tvshow.GET("/top-rated", h.getTopRatedTVShows)
	tvshow.GET("/popular", h.getPopularTVShows)
	tvshow.GET("/genres", h.getTVShowGenres)
}
