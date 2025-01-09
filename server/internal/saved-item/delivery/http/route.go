package http

import "github.com/gin-gonic/gin"

func (h handlerImpl) MapRoutes(r *gin.RouterGroup) {
	r.GET("/saved-item/:user_id", h.getSavedItem)
	r.POST("/saved-item/add", h.addToSavedItem)
	r.POST("/saved-item/remove", h.removeFromSavedItem)
}
