package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	r.GET("/saved-item/:user_id", h.getSavedItem)
	r.PUT("/saved-item/add", h.addToSavedItem)
	r.PUT("/saved-item/remove", h.removeFromSavedItem)
}
