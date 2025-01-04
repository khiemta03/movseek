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
