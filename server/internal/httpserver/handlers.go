package httpserver

import (
	movieHttp "github.com/tmplam/movseek/internal/movie/delivery/http"
	movieRepo "github.com/tmplam/movseek/internal/movie/repository/mongo"
	movieUsecase "github.com/tmplam/movseek/internal/movie/usecase"

	personHttp "github.com/tmplam/movseek/internal/person/delivery/http"
	personRepo "github.com/tmplam/movseek/internal/person/repository/mongo"
	personUsecase "github.com/tmplam/movseek/internal/person/usecase"

	tvShowHttp "github.com/tmplam/movseek/internal/tvshow/delivery/http"
	tvShowRepo "github.com/tmplam/movseek/internal/tvshow/repository/mongo"
	tvShowUsecase "github.com/tmplam/movseek/internal/tvshow/usecase"

	savedItemHttp "github.com/tmplam/movseek/internal/saved-item/delivery/http"
	savedItemRepo "github.com/tmplam/movseek/internal/saved-item/repository/mongo"
	savedItemUsecase "github.com/tmplam/movseek/internal/saved-item/usecase"

	reviewHttp "github.com/tmplam/movseek/internal/review/delivery/http"
	reviewRepo "github.com/tmplam/movseek/internal/review/repository/mongo"
	reviewUsecase "github.com/tmplam/movseek/internal/review/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// Repositories
	movieRepo := movieRepo.NewRepository(srv.db)
	personRepo := personRepo.NewRepository(srv.db)
	tvShowRepo := tvShowRepo.NewRepository(srv.db)
	savedItemRepo := savedItemRepo.NewRepository(srv.db)
	reviewRepo := reviewRepo.NewRepository(srv.db)
	// Usecases
	movieUsecase := movieUsecase.NewUsecase(movieRepo)
	personUsecase := personUsecase.NewUsecase(personRepo)
	tvShowUsecase := tvShowUsecase.NewUsecase(tvShowRepo)
	savedItemUsecase := savedItemUsecase.NewUsecase(savedItemRepo, movieUsecase, tvShowUsecase)
	reviewUsecase := reviewUsecase.NewUsecase(reviewRepo)
	// Handlers
	movieHandler := movieHttp.NewHandler(movieUsecase)
	personHandler := personHttp.NewHandler(personUsecase)
	tvShowHandler := tvShowHttp.NewHandler(tvShowUsecase)
	savedItemHandler := savedItemHttp.NewHandler(savedItemUsecase)
	reviewHandler := reviewHttp.NewHandler(reviewUsecase)

	api := srv.gin.Group("/api/v1")
	movieHandler.MapRoutes(api.Group("/"))
	personHandler.MapRoutes(api.Group("/"))
	tvShowHandler.MapRoutes(api.Group("/"))
	savedItemHandler.MapRoutes(api.Group("/"))
	reviewHandler.MapRoutes(api.Group("/"))
}
