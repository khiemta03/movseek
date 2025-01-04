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
		fmt.Println(err)
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

func (h handlerImpl) searchMovies(c *gin.Context) {
	req, err := h.processSearchMoviesRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	movies, err := h.uc.ListMovies(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, movies)
}
