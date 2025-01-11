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

func (h handlerImpl) processGetOnTheAirTVShowsRequest(c *gin.Context) (getOnTheAirTVShowsRequest, error) {
	var req getOnTheAirTVShowsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getOnTheAirTVShowsRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetAiringTodayTVShowsRequest(c *gin.Context) (getAiringTodayTVShowsRequest, error) {
	var req getAiringTodayTVShowsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getAiringTodayTVShowsRequest{}, err
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
