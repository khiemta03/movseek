package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

func (h handlerImpl) getOneTVShow(c *gin.Context) {
	id, err := h.processGetOneTVShowRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshow, err := h.uc.GetOneTVShow(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, tvshow)
}

func (h handlerImpl) searchTVShows(c *gin.Context) {
	req, err := h.processSearchTVShowsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshows, err := h.uc.ListTVShows(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, tvshows)
}

func (h handlerImpl) getUpcomingTVShows(c *gin.Context) {
	req, err := h.processGetUpcomingTVShowsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshows, err := h.uc.GetUpcomingTVShows(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, tvshows)
}

func (h handlerImpl) getTrendingTVShows(c *gin.Context) {
	req, err := h.processGetTrendingTVShowsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshows, err := h.uc.GetTrendingTVShows(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, tvshows)
}

func (h handlerImpl) getTopRatedTVShows(c *gin.Context) {
	req, err := h.processGetTopRatedTVShowsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshows, err := h.uc.GetTopRatedTVShows(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, tvshows)
}

func (h handlerImpl) getPopularTVShows(c *gin.Context) {
	req, err := h.processGetPopularTVShowsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshows, err := h.uc.GetPopularTVShows(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, tvshows)
}

func (h handlerImpl) getTVShowGenres(c *gin.Context) {
	genres, err := h.uc.GetTVShowGenres(c.Request.Context())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, genres)
}
