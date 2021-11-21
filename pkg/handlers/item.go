package handlers

import (
	"net/http"
	"strconv"

	"github.com/ainurqa95/todo-list-app"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) createItem(context *gin.Context) {
	listId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	var itemInput todo.TodoItem
	if err := context.BindJSON(&itemInput); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := handler.service.ItemManager.CreateItem(itemInput, listId)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"itemId": itemId,
	})
}

type GetAllItemsResponse struct {
	Items []todo.TodoItem `json:"data"`
}

func (handler *Handler) getAllItems(context *gin.Context) {
	userId := getUserId(context)

	listId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	items, err := handler.service.ItemManager.GetAllItems(listId, userId)

	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, GetAllItemsResponse{
		Items: items,
	})

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
