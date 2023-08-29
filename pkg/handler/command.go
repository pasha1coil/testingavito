package handler

import (
	"net/http"

	segment "github.com/pasha1coil/testingavito/pkg/service/enty"

	"github.com/gin-gonic/gin"
)

// @Summary Add-User
// @Tags Add-User
// @Description create account
// @ID Add-User
// @Accept  json
// @Produce  json
// @Param input body segment.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /main/adduser [post]

func (h *Handler) AddUser(c *gin.Context) {
	var input segment.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": id,
	})
}

//  AddSegm godoc
//	@Summary		Add a User
//	@Description	Add User
//	@Accept			json
//	@Produce		json
//	@Param			input	body		segment.Segment true	"User object"
//	@Success		200		{integer}	int
//	@Failure		500		{object}	newErrorResponse
//	@Router			/main/addsegment [post]

func (h *Handler) AddSegm(c *gin.Context) {
	var input segment.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": id,
	})
}

//  DelSegm godoc
//	@Summary		Delete a Slug
//	@Description	Delete Slug
//	@Accept			json
//	@Produce		json
//	@Param			input	body		segment.Segment	true	"Slug object"
//	@Success		200		{integer}	int
//	@Failure		500		{object}	newErrorResponse
//	@Router			/main/delsegment [delete]

func (h *Handler) DelSegm(c *gin.Context) {
	var input segment.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	Status, err := h.services.DelSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": Status,
	})
}

//  InsSegmUsr godoc
//	@Summary		Insert User a Slugs
//	@Description	Insert User a Slugs
//	@Accept			json
//	@Produce		json
//	@Param			input	body		segment.UserSegment	true	"Slug and User objects"
//	@Success		200		{integer}	int
//	@Failure		500		{object}	newErrorResponse
//	@Router			/main/insertsegmentuser [post]

func (h *Handler) InsSegmUsr(c *gin.Context) {
	var input segment.UserSegment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.InsertSemUser(input.NameSegment, input.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//  DelSegmUsr godoc
//	@Summary		Delete User a Slugs
//	@Description	Delete User a Slugs
//	@Accept			json
//	@Produce		json
//	@Param			input	body		segment.UserSegment	true	"Slug and User objects"
//	@Success		200		{integer}	int
//	@Failure		500		{object}	newErrorResponse
//	@Router			/main/delsegmentuser [delete]

func (h *Handler) DelSegmUsr(c *gin.Context) {
	var input segment.UserSegment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.services.DeleteSemUser(input.NameSegment, input.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}

//  GetUsrSegm godoc
//	@Summary		Get User a Slugs
//	@Description	Get User a Slugs
//	@Accept			json
//	@Produce		json
//	@Param			input	body		segment.User	true	"User object"
//	@Success		200		{integer}	int
//	@Failure		500		{object}	newErrorResponse
//	@Router			/main/getusersegment  [post]

func (h *Handler) GetUsrSegm(c *gin.Context) {
	var input segment.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	names, err := h.services.GetActiveSlugs(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(names) != 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"Names": names,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "has no segments or non-existent user entered",
		})
	}

}

//  GetSlugHistoryCsv godoc
//	@Summary		Get User History
//	@Description	Get User History
//	@Accept			json
//	@Produce		json
//	@Param			input	body		segment.GetHistory	true	"User object"
//	@Success		200		{integer}	int
//	@Failure		500		{object}	newErrorResponse
//	@Router			main/history  [get]

func (h *Handler) GetSlugHistoryCsv(c *gin.Context) {
	var input segment.GetHistory
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	csv, err := h.services.GetCsvHistory(input.User_id, input.Start, input.End)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "text/csv; charset=utf-8")

	c.String(http.StatusOK, csv)
}
