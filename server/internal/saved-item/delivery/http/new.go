package http

import (
	"github.com/gin-gonic/gin"
	saved_item "github.com/tmplam/movseek/internal/saved-item"
)

type handler interface {
	MapRoutes(r *gin.RouterGroup)
}

type handlerImpl struct {
	uc saved_item.Usecase
}

func NewHandler(uc saved_item.Usecase) handler {
	return &handlerImpl{uc: uc}
}
