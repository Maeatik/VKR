package handler

import (
	"context"
	v1 "diploma/internal/entity/v1"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	fmt.Println("id: ", userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.User

	user, err = h.service.Service.GetUser(ctx, userId)

	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	fmt.Println("id: ", userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.User

	user, err = h.service.Service.GetUsers(ctx, userId)

	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUsers(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	fmt.Println("id: ", userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.User

	xData := c.Request.Header.Get("X-Data")

	if err := json.Unmarshal([]byte(xData), &user); err != nil {
		h.logger.Info("error while parsing json info about text")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.DeleteUsers(ctx, userId, user.Password)

	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) PutUsers(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	fmt.Println("id: ", userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.UpdateUsers(ctx, userId, user.Name, user.Password)

	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	ctx := context.Background()

	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	var user v1.UserChange

	if err := c.BindJSON(&user); err != nil {
		h.logger.Info("error while parse json")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Service.ChangePassword(ctx, userId, user.Password, user.NewPassword)

	if err != nil {
		fmt.Println(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
