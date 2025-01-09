package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

func (h handlerImpl) getSavedItem(c *gin.Context) {
	req, err := h.processGetSavedItemRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	item, err := h.uc.GetSavedItem(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, item)
}

func (h handlerImpl) addToSavedItem(c *gin.Context) {
	req, err := h.processAddToSavedItemRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.AddToSavedItem(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

func (h handlerImpl) removeFromSavedItem(c *gin.Context) {
	req, err := h.processRemoveFromSavedItemRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.RemoveFromSavedItem(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}
