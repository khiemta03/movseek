package httpserver

import (
	movieHttp "github.com/tmplam/movseek/internal/movie/delivery/http"
	movieRepo "github.com/tmplam/movseek/internal/movie/repository/mongo"
	movieUsecase "github.com/tmplam/movseek/internal/movie/usecase"

	personHttp "github.com/tmplam/movseek/internal/person/delivery/http"
	personRepo "github.com/tmplam/movseek/internal/person/repository/mongo"
	personUsecase "github.com/tmplam/movseek/internal/person/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// Repositories
	movieRepo := movieRepo.NewRepository(srv.db)
	personRepo := personRepo.NewRepository(srv.db)
	// Usecases
	movieUsecase := movieUsecase.NewUsecase(movieRepo)
	personUsecase := personUsecase.NewUsecase(personRepo)
	// Handlers
	movieHandler := movieHttp.NewHandler(movieUsecase)
	personHandler := personHttp.NewHandler(personUsecase)

	api := srv.gin.Group("/api/v1")
	movieHandler.MapRoutes(api.Group("/"))
	personHandler.MapRoutes(api.Group("/"))
}
