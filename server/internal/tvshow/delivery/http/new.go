package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/internal/tvshow"
)

type handler interface {
	MapRoutes(r *gin.RouterGroup)
}

type handlerImpl struct {
	uc tvshow.Usecase
}

func NewHandler(uc tvshow.Usecase) handler {
	return &handlerImpl{uc: uc}
}
