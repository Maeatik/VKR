package handler

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	
	xData := c.Request.Header.Get("X-Data")

	if err := json.Unmarshal([]byte(xData), &textID); err != nil {
		h.logger.Info("error while parsing json info about text")
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

func (h *Handler) DocHandler(c *gin.Context) {
	fmt.Println(1)
	ctx := context.Background()

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}
	var textID v1.TextID
	
	xData := c.Request.Header.Get("X-Data")
	if err := json.Unmarshal([]byte(xData), &textID); err != nil {
		h.logger.Info("error while parsing json info about text")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("text",textID.Id)
	fmt.Println("site",textID.SiteId)
    text, err := h.service.Service.GetMainText(ctx, userId, textID.Id, textID.SiteId)
	if err != nil {
		h.logger.Info("error while getting text")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

    file := []byte(text.Text)
	fmt.Println(file)
   	if err := ioutil.WriteFile("document.doc", file, 0644); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать документ"})
        return
    }

    c.Writer.Header().Set("Content-Disposition", "attachment; filename=document.doc")
    c.Writer.Header().Set("Content-Type", "application/msword")
    c.File("document.doc")
}

