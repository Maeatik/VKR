package handler

import (
	//"net/http"
	"context"
	v1 "diploma/internal/entity/v1"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) GetSite(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}
	var siteID v1.SiteID

	if err := c.BindJSON(&siteID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	site, err := h.service.Service.GetSite(ctx, userId, siteID.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, site)
}
func (h *Handler) PostSite(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var site v1.Site

	if err := c.BindJSON(&site); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	siteID, err := h.service.Service.PostSite(ctx, userId, site.Url, site.Tag)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, siteID)
}

func (h *Handler) DeleteSite(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}
	var site v1.SiteID

	xData := c.Request.Header.Get("X-Data")

	if err := json.Unmarshal([]byte(xData), &site); err != nil {
		h.logger.Info("error while parsing json info about text")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.DeleteSite(ctx, userId, site.Id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, site)
}

func (h *Handler) GetListSites(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	site, err := h.service.Service.GetListSites(ctx, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, site)
}

func (h *Handler) ParseSite(c *gin.Context) {
	h.logger.Info(">>Start Parsing")
	ctx := context.Background()

	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var site v1.Site
	if err := c.BindJSON(&site); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.ParseSite(ctx, userId, site.Url, site.Tag)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userId)
	h.logger.Info(">>End Parsing")
}
