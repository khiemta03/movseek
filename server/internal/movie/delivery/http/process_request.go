package http

import "github.com/gin-gonic/gin"

func (h handler) processGetOneMovieRequest(c *gin.Context) (string, error) {
	id := c.Param("id")

	return id, nil
}

func (h handler) processGetMovieCreditsRequest(c *gin.Context) (string, error) {
	id := c.Param("id")

	return id, nil
}
