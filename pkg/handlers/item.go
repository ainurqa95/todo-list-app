package handlers

import (
	"net/http"
	"strconv"

	"github.com/ainurqa95/todo-list-app"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) createItem(context *gin.Context) {
	listId, err := strconv.Atoi(context.Param("id"))
	userId := getUserId(context)

	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	var itemInput todo.TodoItem
	if err := context.BindJSON(&itemInput); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := handler.service.ItemManager.CreateItem(itemInput, userId, listId)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"itemId": itemId,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
