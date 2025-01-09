package http

import (
	"github.com/gin-gonic/gin"
)

func (h handlerImpl) processAddCommentRequest(c *gin.Context) (addCommentRequest, error) {
	var req addCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return addCommentRequest{}, err
	}

	if err := req.validate(); err != nil {
		return addCommentRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processAddRatingRequest(c *gin.Context) (addRatingRequest, error) {
	var req addRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return addRatingRequest{}, err
	}

	if err := req.validate(); err != nil {
		return addRatingRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processUpdateCommentRequest(c *gin.Context) (updateCommentRequest, error) {
	var req updateCommentRequest
	if err := c.ShouldBindUri(&req); err != nil {
		return updateCommentRequest{}, err
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		return updateCommentRequest{}, err
	}

	if err := req.validate(); err != nil {
		return updateCommentRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processUpdateRatingRequest(c *gin.Context) (updateRatingRequest, error) {
	var req updateRatingRequest
	if err := c.ShouldBindUri(&req); err != nil {
		return updateRatingRequest{}, err
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		return updateRatingRequest{}, err
	}

	if err := req.validate(); err != nil {
		return updateRatingRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processDeleteCommentRequest(c *gin.Context) (deleteCommentRequest, error) {
	var req deleteCommentRequest
	if err := c.ShouldBindUri(&req); err != nil {
		return deleteCommentRequest{}, err
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		return deleteCommentRequest{}, err
	}

	if err := req.validate(); err != nil {
		return deleteCommentRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processDeleteRatingRequest(c *gin.Context) (deleteRatingRequest, error) {
	var req deleteRatingRequest
	if err := c.ShouldBindUri(&req); err != nil {
		return deleteRatingRequest{}, err
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		return deleteRatingRequest{}, err
	}

	if err := req.validate(); err != nil {
		return deleteRatingRequest{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetCommentsByMediaIDRequest(c *gin.Context) (getCommentsRequestByMediaID, error) {
	var req getCommentsRequestByMediaID
	if err := c.ShouldBindUri(&req); err != nil {
		return getCommentsRequestByMediaID{}, err
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		return getCommentsRequestByMediaID{}, err
	}

	if err := req.validate(); err != nil {
		return getCommentsRequestByMediaID{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetRatingsByMediaIDRequest(c *gin.Context) (getRatingsRequestByMediaID, error) {
	var req getRatingsRequestByMediaID
	if err := c.ShouldBindUri(&req); err != nil {
		return getRatingsRequestByMediaID{}, err
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		return getRatingsRequestByMediaID{}, err
	}

	if err := req.validate(); err != nil {
		return getRatingsRequestByMediaID{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetRatingsByUserIDRequest(c *gin.Context) (getRatingsRequestByUserID, error) {
	var req getRatingsRequestByUserID
	if err := c.ShouldBindUri(&req); err != nil {
		return getRatingsRequestByUserID{}, err
	}

	return req, nil
}

func (h handlerImpl) processGetCommentsByUserIDRequest(c *gin.Context) (getCommentsRequestByUserID, error) {
	var req getCommentsRequestByUserID
	if err := c.ShouldBindUri(&req); err != nil {
		return getCommentsRequestByUserID{}, err
	}

	return req, nil
}
