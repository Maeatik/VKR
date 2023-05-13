package handler

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetText(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}
	var textID v1.TextID

	if err := c.BindJSON(&textID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	text, err := h.service.Service.GetMainText(ctx, userId, textID.Id, textID.SiteId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, text)
}
func (h *Handler) PostText(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var text v1.MainText
	if err := c.BindJSON(&text); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	text.Date = time.Now()

	err = h.service.Service.PostMainText(ctx, userId, text.SiteID, text.Text)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, text)
}

func (h *Handler) PutText(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var text v1.MainText

	if err := c.BindJSON(&text); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.UpdateMainText(ctx, userId, text.Id, text.SiteID, text.Text)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "updated")
}
