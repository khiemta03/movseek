package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handlerImpl) processGetOneMovieRequest(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (h handlerImpl) processGetMovieCreditsRequest(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (h handlerImpl) processGetMovieVideosRequest(c *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

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

func (h handlerImpl) processGetTopRatedMoviesRequest(c *gin.Context) (getTopRatedMoviesRequest, error) {
	var req getTopRatedMoviesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getTopRatedMoviesRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetPopularMoviesRequest(c *gin.Context) (getPopularMoviesRequest, error) {
	var req getPopularMoviesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getPopularMoviesRequest{}, err
	}

	return req, nil
}
