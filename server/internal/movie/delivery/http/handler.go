package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

func (h *handler) getOneMovie(c *gin.Context) {
	id, err := h.processGetOneMovieRequest(c)
	if err != nil {
		response.Forbidden(c)
		return
	}

	movie, err := h.uc.GetOneMovie(c.Request.Context(), id)
	if err != nil {
		fmt.Println(err)
		response.Forbidden(c)
		return
	}

	response.OK(c, movie)
}

func (h *handler) getMovieCredits(c *gin.Context) {
	id, err := h.processGetMovieCreditsRequest(c)
	if err != nil {
		response.Forbidden(c)
		return
	}

	credits, err := h.uc.GetMovieCredits(c.Request.Context(), id)
	if err != nil {
		response.Forbidden(c)
		return
	}

	response.OK(c, credits)
}
