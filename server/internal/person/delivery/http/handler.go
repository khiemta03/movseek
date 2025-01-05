package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

func (h handlerImpl) getOnePerson(c *gin.Context) {
	id, err := h.processGetOnePersonRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	person, err := h.uc.GetOnePerson(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, person)
}

func (h handlerImpl) searchPeople(c *gin.Context) {
	req, err := h.processSearchPeopleRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	people, err := h.uc.ListPeople(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, people)
}

func (h handlerImpl) getTrendingPeople(c *gin.Context) {
	req, err := h.processGetTrendingPeopleRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	people, err := h.uc.GetTrendingPeople(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, people)
}

func (h handlerImpl) getPopularPeople(c *gin.Context) {
	req, err := h.processGetPopularPeopleRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	people, err := h.uc.GetPopularPeople(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, people)
}
