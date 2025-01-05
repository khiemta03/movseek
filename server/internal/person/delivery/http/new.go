package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/internal/person"
)

type handler interface {
	MapRoutes(r *gin.RouterGroup)
}

type handlerImpl struct {
	uc person.Usecase
}

func NewHandler(uc person.Usecase) handler {
	return &handlerImpl{uc: uc}
}
