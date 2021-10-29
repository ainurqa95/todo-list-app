package handler

import (
	"net/http"

	"github.com/ainurqa95/todo-list-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(context *gin.Context) {
	var input todo.User
	if err := context.BindJSON(&input); err != nil {
		newErrorMessage(context, http.StatusBadRequest, err.Error())
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorMessage(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
