package httpserver

import (
	"github.com/tmplam/movseek/internal/movie/delivery/http"
	movieRepo "github.com/tmplam/movseek/internal/movie/repository/mongo"
	movieUsecase "github.com/tmplam/movseek/internal/movie/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// Repositories
	movieRepo := movieRepo.NewRepository(srv.db)

	// Usecases
	movieUsecase := movieUsecase.NewUsecase(movieRepo)
	// Handlers
	movieHandler := http.NewHandler(movieUsecase)

	api := srv.gin.Group("/api/v1")
	movieHandler.MapRoutes(api.Group("/movies"))
}
