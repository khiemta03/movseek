package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tmplam/movseek/pkg/response"
)

// @Summary Add comment
// @Description Add comment
// @Tags review
// @Accept json
// @Produce json
// @Param addCommentRequest body addCommentRequest true "Add comment request"
// @Router /comments [post]
func (h handlerImpl) addComment(c *gin.Context) {
	req, err := h.processAddCommentRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	cm, err := h.uc.AddComment(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, cm)
}

// @Summary Update comment
// @Description Update comment
// @Tags review
// @Accept json
// @Produce json
// @Param updateCommentRequest body updateCommentRequest true "Update comment request"
// @Router /comments/{user_id} [put]
func (h handlerImpl) updateComment(c *gin.Context) {
	req, err := h.processUpdateCommentRequest(c)
	if err != nil {
		fmt.Println("1: ", err)
		response.BadRequest(c)
		return
	}

	err = h.uc.UpdateComment(c.Request.Context(), req.toInput())
	if err != nil {
		fmt.Println("2: ", err)
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

// @Summary Delete comment
// @Description Delete comment
// @Tags review
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param media_id query int true "Media ID"
// @Param type query string true "Type"
// @Router /comments/{user_id} [delete]
func (h handlerImpl) deleteComment(c *gin.Context) {
	req, err := h.processDeleteCommentRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.DeleteComment(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

// @Summary Add rating
// @Description Add rating
// @Tags review
// @Accept json
// @Produce json
// @Param addRatingRequest body addRatingRequest true "Add rating request"
// @Router /ratings [post]
func (h handlerImpl) addRating(c *gin.Context) {
	req, err := h.processAddRatingRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	rt, err := h.uc.AddRating(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, rt)
}

// @Summary Update rating
// @Description Update rating
// @Tags review
// @Accept json
// @Produce json
// @Param updateRatingRequest body updateRatingRequest true "Update rating request"
// @Router /ratings/{user_id} [put]
func (h handlerImpl) updateRating(c *gin.Context) {
	req, err := h.processUpdateRatingRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.UpdateRating(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

// @Summary Delete rating
// @Description Delete rating
// @Tags review
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param media_id query int true "Media ID"
// @Param type query string true "Type"
// @Router /ratings/{user_id} [delete]
func (h handlerImpl) deleteRating(c *gin.Context) {
	req, err := h.processDeleteRatingRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	err = h.uc.DeleteRating(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, nil)
}

// @Summary Get comments by media
// @Description Get comments by media
// @Tags review
// @Accept json
// @Produce json
// @Param media_id path int true "Media ID"
// @Router /comments/media/{media_id} [get]
func (h handlerImpl) getCommentsByMedia(c *gin.Context) {
	req, err := h.processGetCommentsByMediaIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	cm, err := h.uc.GetCommentsByMedia(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.getCommentsByMediaResponse(cm))
}

// @Summary Get ratings by media
// @Description Get ratings by media
// @Tags review
// @Accept json
// @Produce json
// @Param media_id path int true "Media ID"
// @Router /ratings/media/{media_id} [get]
func (h handlerImpl) getRatingsByMedia(c *gin.Context) {
	req, err := h.processGetRatingsByMediaIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	rt, err := h.uc.GetRatingsByMedia(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.getRatingsByMediaResponse(rt))
}

// @Summary Get comments by user
// @Description Get comments by user
// @Tags review
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Router /comments/user/{user_id} [get]
func (h handlerImpl) getCommentsByUser(c *gin.Context) {
	req, err := h.processGetCommentsByUserIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	cm, err := h.uc.GetCommentsByUserID(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.getCommentsByUserResponse(cm))
}

// @Summary Get ratings by user
// @Description Get ratings by user
// @Tags review
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Router /ratings/user/{user_id} [get]
func (h handlerImpl) getRatingsByUser(c *gin.Context) {
	req, err := h.processGetRatingsByUserIDRequest(c)
	if err != nil {
		response.BadRequest(c)
		return
	}

	rt, err := h.uc.GetRatingsByUserID(c.Request.Context(), req.toInput())
	if err != nil {
		response.BadRequest(c)
		return
	}

	response.OK(c, h.getRatingsByUserResponse(rt))
}
