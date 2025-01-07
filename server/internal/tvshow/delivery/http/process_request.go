package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handlerImpl) processGetOneTVShowRequest(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (h handlerImpl) processSearchTVShowsRequest(c *gin.Context) (searchTVShowsRequest, error) {
	var req searchTVShowsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return searchTVShowsRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetUpcomingTVShowsRequest(c *gin.Context) (getUpcomingTVShowsRequest, error) {
	var req getUpcomingTVShowsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getUpcomingTVShowsRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetTrendingTVShowsRequest(c *gin.Context) (getTrendingTVShowsRequest, error) {
	var req getTrendingTVShowsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getTrendingTVShowsRequest{}, err
	}

	if err := c.ShouldBindUri(&req); err != nil {
		return getTrendingTVShowsRequest{}, err
	}

	if err := req.validate(); err != nil {
		return getTrendingTVShowsRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetTopRatedTVShowsRequest(c *gin.Context) (getTopRatedTVShowsRequest, error) {
	var req getTopRatedTVShowsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getTopRatedTVShowsRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetPopularTVShowsRequest(c *gin.Context) (getPopularTVShowsRequest, error) {
	var req getPopularTVShowsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getPopularTVShowsRequest{}, err
	}

	return req, nil
}
