package http

import (
	"github.com/gin-gonic/gin"
)

func (h handlerImpl) processGetOneMovieRequest(c *gin.Context) (string, error) {
	id := c.Param("id")

	return id, nil
}

func (h handlerImpl) processGetMovieCreditsRequest(c *gin.Context) (string, error) {
	id := c.Param("id")

	return id, nil
}

func (h handlerImpl) processSearchMoviesRequest(c *gin.Context) (searchMoviesRequest, error) {
	var req searchMoviesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return searchMoviesRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetUpcomingMoviesRequest(c *gin.Context) (getUpcomingMoviesRequest, error) {
	var req getUpcomingMoviesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getUpcomingMoviesRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetTrendingMoviesRequest(c *gin.Context) (getTrendingMoviesRequest, error) {
	var req getTrendingMoviesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getTrendingMoviesRequest{}, err
	}

	if err := c.ShouldBindUri(&req); err != nil {
		return getTrendingMoviesRequest{}, err
	}

	if err := req.validate(); err != nil {
		return getTrendingMoviesRequest{}, err
	}

	return req, nil
}
