package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

func (h handlerImpl) addComment(c *gin.Context) {
	req, err := h.processAddCommentRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	cm, err := h.uc.AddComment(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, cm)
}

func (h handlerImpl) updateComment(c *gin.Context) {
	req, err := h.processUpdateCommentRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.UpdateComment(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

func (h handlerImpl) deleteComment(c *gin.Context) {
	req, err := h.processDeleteCommentRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.DeleteComment(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

func (h handlerImpl) addRating(c *gin.Context) {
	req, err := h.processAddRatingRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	rt, err := h.uc.AddRating(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, rt)
}

func (h handlerImpl) updateRating(c *gin.Context) {
	req, err := h.processUpdateRatingRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.UpdateRating(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

func (h handlerImpl) deleteRating(c *gin.Context) {
	req, err := h.processDeleteRatingRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.DeleteRating(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

func (h handlerImpl) getCommentsByMedia(c *gin.Context) {
	req, err := h.processGetCommentsByMediaIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	cm, err := h.uc.GetCommentsByMedia(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, cm)
}

func (h handlerImpl) getRatingsByMedia(c *gin.Context) {
	req, err := h.processGetRatingsByMediaIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	rt, err := h.uc.GetRatingsByMedia(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, rt)
}

func (h handlerImpl) getCommentsByUser(c *gin.Context) {
	req, err := h.processGetCommentsByUserIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	cm, err := h.uc.GetCommentsByUserID(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, cm)
}

func (h handlerImpl) getRatingsByUser(c *gin.Context) {
	req, err := h.processGetRatingsByUserIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	rt, err := h.uc.GetRatingsByUserID(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, rt)
}
