package handler

import (
	"net/http"
	segment "testingavito"

	"github.com/gin-gonic/gin"
)

// AddUser	godoc
//
// @Summary	AddUser
// @Description	create user
// @Accept	json
// @Produce	json
// @Param	input 	body	segment.User 	true	"user info"
// @Success	200	{integer}	integer 1
// @Failure 500 {object} errorResponse
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
		"UserNumber": id,
	})
}

// AddSegm godoc
//
// @Summary AddSegm
// @Description create slug
// @Accept  json
// @Produce  json
// @Param input body segment.Segment true "slug info"
// @Success 200 {integer} integer 1
// @Failure 500 {object} errorResponse
// @Router /main/addsegment [post]

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
		"SlagName": id,
	})
}

// DelSegm godoc
//
// @Summary DelSegm
// @Description delete slug
// @Accept  json
// @Produce  json
// @Param input body segment.Segment true "slug info"
// @Success 200 {integer} integer 1
// @Failure 500 {object} errorResponse
// @Router /main/delsegment [delete]

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

// InsSegmUsr godoc
//
// @Summary InsSegmUsr
// @Description insert slug to user
// @Accept  json
// @Produce  json
// @Param input body segment.UserSegment true "slug and user info"
// @Success 200 {integer} integer 1
// @Failure 500 {object} errorResponse
// @Router /main/insertsegmentuser [post]

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

// DelSegmUsr godoc
//
// @Summary DelSegmUsr
// @Description delete slug to user
// @Accept  json
// @Produce  json
// @Param input body segment.UserSegment true "slug and user info"
// @Success 200 {integer} integer 1
// @Failure 500 {object} errorResponse
// @Router /main/delsegmentuser [delete]

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

// GetUsrSegm godoc
//
// @Summary GetUsrSegm
// @Description get user slugs
// @Accept  json
// @Produce  json
// @Param input body segment.User true "user info"
// @Success 200 {integer} integer 1
// @Failure 500 {object} errorResponse
// @Router /main/getusersegment [post]

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
			"Names": "не имеет сегментов",
		})
	}

}

// GetSlugHistoryCsv godoc
//
// @Summary GetSlugHistoryCsv
// @Description get csv otchet for user slugs
// @Accept  json
// @Produce  json
// @Param input body segment.GetHistory true "user info, start and end date YY/MM"
// @Success 200 {integer} integer 1
// @Failure 500 {object} errorResponse
// @Router /main/history [get]

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
