package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

// @Summary Get one movie
// @Description Get one movie
// @Tags movie
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} models.Movie
// @Router /movie/{id} [get]
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

// @Summary Get movie credits
// @Description Get movie credits
// @Tags movie
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} models.MovieCredits
// @Router /movie/{id}/credits [get]
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

// @Summary Get movie videos
// @Description Get movie videos
// @Tags movie
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} []models.Trailer
// @Router /movie/{id}/videos [get]
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

// @Summary Get movie keywords
// @Description Get movie keywords
// @Tags movie
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} []models.KeyWord
// @Router /movie/{id}/keywords [get]
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

// @Summary Search movies
// @Description Search movies
// @Tags movie
// @Accept json
// @Produce json
// @Param query query string false "Query"
// @Param page query int true "Page"
// @Param per_page query int true "Per page"
// @Success 200 {object} searchMoviesResponse
// @Router /search/movie [get]
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

// @Summary Get upcoming movies
// @Description Get upcoming movies
// @Tags movie
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param per_page query int true "Per page"
// @Success 200 {object} getSummaryMoviesResponse
// @Router /movie/upcoming [get]
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

// @Summary Get trending movies
// @Description Get trending movies
// @Tags movie
// @Accept json
// @Produce json
// @Param type path string true "Type"
// @Param page query int true "Page"
// @Param per_page query int true "Per page"
// @Success 200 {object} getSummaryMoviesResponse
// @Router /movie/trending/{type} [get]
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

// @Summary Get top rated movies
// @Description Get top rated movies
// @Tags movie
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param per_page query int true "Per page"
// @Success 200 {object} getSummaryMoviesResponse
// @Router /movie/top-rated [get]
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

// @Summary Get popular movies
// @Description Get popular movies
// @Tags movie
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param per_page query int true "Per page"
// @Success 200 {object} getSummaryMoviesResponse
// @Router /movie/popular [get]
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

// @Summary Get now playing movies
// @Description Get now playing movies
// @Tags movie
// @Accept json
// @Produce json
// @Param page query int true "Page"
// @Param per_page query int true "Per page"
// @Success 200 {object} getSummaryMoviesResponse
// @Router /movie/now-playing [get]
func (h handlerImpl) getNowPlayingMovies(c *gin.Context) {
	req, err := h.processGetNowPlayingMoviesRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	movies, err := h.uc.GetNowPlayingMovies(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.buildGetSummaryMoviesResponse(movies.Movies, movies.Pagination))
}

// @Summary Get movie genres
// @Description Get movie genres
// @Tags movie
// @Accept json
// @Produce json
// @Success 200 {object} movie.GetMovieGenresOutput
// @Router /movie/genres [get]
func (h handlerImpl) getMovieGenres(c *gin.Context) {
	genres, err := h.uc.GetMovieGenres(c.Request.Context())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, genres)
}

// @Summary Get lastest trailer
// @Description Get lastest trailer
// @Tags movie
// @Accept json
// @Produce json
// @Success 200 {object} []models.Trailer
// @Router /movie/trailer/latest [get]
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
