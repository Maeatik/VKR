package handler

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}


func newErrorResponse(c *gin.Context, statusCode int, msg string)  {
	c.AbortWithStatusJSON(statusCode, Error{msg})
}