package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handlerImpl) processGetOnePersonRequest(c *gin.Context) (int64, error) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, err
	}

	return idInt, nil
}

func (h handlerImpl) processSearchPeopleRequest(c *gin.Context) (searchPeopleRequest, error) {
	var req searchPeopleRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return searchPeopleRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetTrendingPeopleRequest(c *gin.Context) (getTrendingPeopleRequest, error) {
	var req getTrendingPeopleRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getTrendingPeopleRequest{}, err
	}

	if err := c.ShouldBindUri(&req); err != nil {
		return getTrendingPeopleRequest{}, err
	}

	if err := req.validate(); err != nil {
		return getTrendingPeopleRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetPopularPeopleRequest(c *gin.Context) (getPopularPeopleRequest, error) {
	var req getPopularPeopleRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return getPopularPeopleRequest{}, err
	}

	return req, nil
}
