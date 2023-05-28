package handler

import (
	v1 "diploma/internal/entity/v1"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// регистрация
func (h *Handler) Register (c *gin.Context){
	var input v1.User
	fmt.Println(1)
	if err :=c.BindJSON(&input); err != nil{
		h.logger.Info("error while parse json request")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(2)
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if id == 0{
		newErrorResponse(c, http.StatusNotAcceptable, "Имя аккаунта уже занято")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type regInput struct {
	Name 		string 	`json:"name" binding:"required"`
	Password 	string 	`json:"password" binding:"required"`
}

func (h *Handler) Login (c *gin.Context)  {
	var input regInput
	
	if err :=c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Name, input.Password)

	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": token,
	})
}