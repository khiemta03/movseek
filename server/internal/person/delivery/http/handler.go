package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

// @Summary Get one person
// @Description Get one person
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} models.Person
// @Router /person/{id} [get]
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

// @Summary Search people
// @Description Search people
// @Tags person
// @Accept json
// @Produce json
// @Param query query string true "Query"
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Success 200 {object} getListPeopleResponse
// @Router /search/person [get]
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

	response.OK(c, h.toResponse(people))
}

// @Summary Get trending people
// @Description Get trending people
// @Tags person
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Success 200 {object} getListPeopleSummaryResponse
// @Router /person/trending/{type} [get]
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

	response.OK(c, h.toSummaryResponse(people.People, people.Pagination))
}

// @Summary Get popular people
// @Description Get popular people
// @Tags person
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Success 200 {object} getListPeopleSummaryResponse
// @Router /person/popular [get]
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

	response.OK(c, h.toSummaryResponse(people.People, people.Pagination))
}

// @Summary Get person credits
// @Description Get person credits
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} getListPeopleSummaryResponse
// @Router /person/{id}/credits/movie [get]
func (h handlerImpl) getPersonCredits(c *gin.Context) {
	id, err := h.processGetPersonCreditsRequest(c)
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

// @Summary Get person TV credits
// @Description Get person TV credits
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} getListPeopleSummaryResponse
// @Router /person/{id}/credits/tv [get]
func (h handlerImpl) getTVCredits(c *gin.Context) {
	id, err := h.processGetTVCreditsRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	credits, err := h.uc.GetTVCredits(c.Request.Context(), id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, credits)
}
