package handler

import (
	v1 "diploma/internal/entity/v1"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user1 []v1.User

func (h *Handler) GetUsers (c *gin.Context){

	userId, err := getUserId(c)

	fmt.Println("id: ",userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.User

	if err := c.BindJSON(&user); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.service.Service.GetUsers(user.Name, user.Password)

	if err != nil{
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) PostUsers (c *gin.Context){

}
func (h *Handler) DeleteUsers (c *gin.Context){
	userId, err := getUserId(c)

	fmt.Println("id: ",userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.User

	if err := c.BindJSON(&user); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.DeleteUsers(userId)

	if err != nil{
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) PutUsers (c *gin.Context){
	userId, err := getUserId(c)

	fmt.Println("id: ",userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.User

	if err := c.BindJSON(&user); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.UpdateUsers(userId, user.Name, user.Password)

	if err != nil{
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}