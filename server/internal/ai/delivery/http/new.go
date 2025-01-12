package http

import (
	"github.com/gin-gonic/gin"
)

type handler interface {
	MapRoutes(r *gin.RouterGroup)
}

type handlerImpl struct {
}

func NewHandler() handler {
	return &handlerImpl{}
}
