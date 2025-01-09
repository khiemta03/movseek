package http

import (
	"github.com/gin-gonic/gin"
	review "github.com/tmplam/movseek/internal/review"
)

type handler interface {
	MapRoutes(r *gin.RouterGroup)
}

type handlerImpl struct {
	uc review.Usecase
}

func NewHandler(uc review.Usecase) handler {
	return &handlerImpl{uc: uc}
}
