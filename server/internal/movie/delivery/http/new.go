package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/internal/movie"
)

type handler interface {
	MapRoutes(r *gin.RouterGroup)
}

type handlerImpl struct {
	uc movie.Usecase
}

func NewHandler(uc movie.Usecase) handler {
	return &handlerImpl{uc: uc}
}
