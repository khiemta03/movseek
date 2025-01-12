package http

import (
	"github.com/gin-gonic/gin"
)

func (h handlerImpl) processRetrieverRequest(c *gin.Context) (retrieverRequest, error) {
	var req retrieverRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return retrieverRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processNavigateRequest(c *gin.Context) (navigateRequest, error) {
	var req navigateRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return navigateRequest{}, err
	}

	return req, nil
}
