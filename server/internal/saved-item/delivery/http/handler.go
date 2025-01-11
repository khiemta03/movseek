package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

// @Summary Get saved item
// @Description Get saved item
// @Tags saved-item
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.SavedItems
// @Router /saved-item/{user_id} [get]
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

// @Summary Add to saved item
// @Description Add to saved item
// @Tags saved-item
// @Accept json
// @Produce json
// @Param addToSavedItemRequest body addToSavedItemRequest true "Add to saved item request"
// @Router /saved-item/add [put]
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

// @Summary Add to saved item
// @Description Add to saved item
// @Tags saved-item
// @Accept json
// @Produce json
// @Param removeFromSavedItemRequest body removeFromSavedItemRequest true "Remove from saved item request"
// @Router /saved-item/remove [put]
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
