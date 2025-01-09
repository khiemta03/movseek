package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

func (h handlerImpl) getOneMovie(c *gin.Context) {
	id, err := h.processGetOneMovieRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	movie, err := h.uc.GetOneMovie(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, movie)
}

func (h handlerImpl) getMovieCredits(c *gin.Context) {
	id, err := h.processGetMovieCreditsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	credits, err := h.uc.GetMovieCredits(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, credits)
}

func (h handlerImpl) getMovieVideos(c *gin.Context) {
	id, err := h.processGetMovieVideosRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	videos, err := h.uc.GetMovieVideos(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, videos)
}

func (h handlerImpl) getMovieKeywords(c *gin.Context) {
	id, err := h.processGetMovieKeywordsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	keywords, err := h.uc.GetMovieKeywords(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, keywords)
}

func (h handlerImpl) searchMovies(c *gin.Context) {
	req, err := h.processSearchMoviesRequest(c)
	if err != nil {
		fmt.Println(err)
		response.BadRequest(c)
		return
	}

	movies, err := h.uc.ListMovies(c.Request.Context(), req.toInput())
	if err != nil {
		fmt.Println(err)
		response.BadRequest(c)
		return
	}

	response.OK(c, h.buildSearchMoviesResponse(movies))
}

func (h handlerImpl) getUpcomingMovies(c *gin.Context) {
	req, err := h.processGetUpcomingMoviesRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	movies, err := h.uc.GetUpcomingMovies(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.buildGetSummaryMoviesResponse(movies.Movies, movies.Pagination))
}

func (h handlerImpl) getTrendingMovies(c *gin.Context) {
	req, err := h.processGetTrendingMoviesRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	movies, err := h.uc.GetTrendingMovies(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.buildGetSummaryMoviesResponse(movies.Movies, movies.Pagination))
}

func (h handlerImpl) getTopRatedMovies(c *gin.Context) {
	req, err := h.processGetTopRatedMoviesRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	movies, err := h.uc.GetTopRatedMovies(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.buildGetSummaryMoviesResponse(movies.Movies, movies.Pagination))
}

func (h handlerImpl) getPopularMovies(c *gin.Context) {
	req, err := h.processGetPopularMoviesRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	movies, err := h.uc.GetPopularMovies(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.buildGetSummaryMoviesResponse(movies.Movies, movies.Pagination))
}

func (h handlerImpl) getMovieGenres(c *gin.Context) {
	genres, err := h.uc.GetMovieGenres(c.Request.Context())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, genres)
}

func (h handlerImpl) getLastestTrailer(c *gin.Context) {
	req, err := h.processGetLastestTrailerRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	trailers, err := h.uc.GetLastestTrailer(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, trailers)
}
