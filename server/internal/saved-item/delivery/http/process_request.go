package http

import (
	"github.com/gin-gonic/gin"
)

func (h handlerImpl) processGetSavedItemRequest(c *gin.Context) (getSavedItemRequest, error) {
	var req getSavedItemRequest
	if err := c.ShouldBindUri(&req); err != nil {
		return getSavedItemRequest{}, err
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		return getSavedItemRequest{}, err
	}

	if err := req.validate(); err != nil {
		return getSavedItemRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processAddToSavedItemRequest(c *gin.Context) (addToSavedItemRequest, error) {
	var req addToSavedItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return addToSavedItemRequest{}, err
	}

	if err := req.validate(); err != nil {
		return addToSavedItemRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processRemoveFromSavedItemRequest(c *gin.Context) (removeFromSavedItemRequest, error) {
	var req removeFromSavedItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return removeFromSavedItemRequest{}, err
	}

	if err := req.validate(); err != nil {
		return removeFromSavedItemRequest{}, err
	}

	return req, nil
}
