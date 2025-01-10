package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

// @Summary Get one tv show
// @Description Get one tv show
// @Tags tv show
// @Accept json
// @Produce json
// @Param id path int true "TV Show ID"
// @Success 200 {object} models.TVShow
// @Router /tvshow/{id} [get]
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

// @Summary Search tv shows
// @Description Search tv shows
// @Tags tv show
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param page query int true "Page"
// @Param pageSize query int true "Page size"
// @Success 200 {object} getListTVShowsResponse
// @Router /tvshow/search [get]
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

	response.OK(c, h.toResponse(tvshows))
}

// @Summary Get upcoming tv shows
// @Description Get upcoming tv shows
// @Tags tv show
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param pageSize query int true "Page size"
// @Success 200 {object} getListTVShowSummaryResponse
// @Router /tvshow/upcoming [get]
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

	response.OK(c, h.toSummaryResponse(tvshows.TVShows, tvshows.Pagination))
}

// @Summary Get on the air tv shows
// @Description Get on the air tv shows
// @Tags tv show
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param pageSize query int true "Page size"
// @Success 200 {object} getListTVShowSummaryResponse
// @Router /tvshow/on-the-air [get]
func (h handlerImpl) getOnTheAirTVShows(c *gin.Context) {
	req, err := h.processGetOnTheAirTVShowsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshows, err := h.uc.GetOnTheAirTVShows(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.toSummaryResponse(tvshows.TVShows, tvshows.Pagination))
}

// @Summary Get airing today tv shows
// @Description Get airing today tv shows
// @Tags tv show
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param pageSize query int true "Page size"
// @Success 200 {object} getListTVShowSummaryResponse
// @Router /tvshow/airing-today [get]
func (h handlerImpl) getAiringTodayTVShows(c *gin.Context) {
	req, err := h.processGetAiringTodayTVShowsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	tvshows, err := h.uc.GetAiringTodayTVShows(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.toSummaryResponse(tvshows.TVShows, tvshows.Pagination))
}

// @Summary Get top rated tv shows
// @Description Get top rated tv shows
// @Tags tv show
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param pageSize query int true "Page size"
// @Success 200 {object} getListTVShowSummaryResponse
// @Router /tvshow/top-rated [get]
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

	response.OK(c, h.toSummaryResponse(tvshows.TVShows, tvshows.Pagination))
}

// @Summary Get popular tv shows
// @Description Get popular tv shows
// @Tags tv show
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param pageSize query int true "Page size"
// @Success 200 {object} getListTVShowSummaryResponse
// @Router /tvshow/popular [get]
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

	response.OK(c, h.toSummaryResponse(tvshows.TVShows, tvshows.Pagination))
}

// @Summary Get tv show genres
// @Description Get tv show genres
// @Tags tv show
// @Accept json
// @Produce json
// @Success 200 {object} []models.TVGenre
// @Router /tvshow/genres [get]
func (h handlerImpl) getTVShowGenres(c *gin.Context) {
	genres, err := h.uc.GetTVShowGenres(c.Request.Context())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, genres)
}
