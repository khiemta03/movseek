package httpserver

import (
	"github.com/tmplam/movseek/pkg/mongo"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	gin  *gin.Engine
	port int
	db   mongo.Database
}

type Config struct {
	Port int
	DB   mongo.Database
}

func New(cfg Config) *HTTPServer {
	return &HTTPServer{
		gin:  gin.Default(),
		port: cfg.Port,
		db:   cfg.DB,
	}
}
