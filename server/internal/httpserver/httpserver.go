package httpserver

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Run starts the HTTP server.
func (srv HTTPServer) Run() {
	srv.mapHandlers()

	go func() {
		srv.gin.Run(fmt.Sprintf(":%d", srv.port))
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
