package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx = "id"
)


func (h *Handler) userIdentity(c *gin.Context){
	header := c.GetHeader(authHeader)
	if header == ""{
		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2{
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil{
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}