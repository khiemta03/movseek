package httpserver

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run starts the HTTP server.
func (srv HTTPServer) Run() {
	srv.mapHandlers()
	srv.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	go func() {
		srv.gin.Run(fmt.Sprintf(":%d", srv.port))
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
