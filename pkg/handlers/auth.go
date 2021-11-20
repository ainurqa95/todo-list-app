package handlers

import (
	"net/http"

	"github.com/ainurqa95/todo-list-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(context *gin.Context) {
	var input todo.User
	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(context *gin.Context) {
	var input signInInput

	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, 402, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
