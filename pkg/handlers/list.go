package handlers

import (
	"net/http"
	"strconv"

	"github.com/ainurqa95/todo-list-app"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) createList(context *gin.Context) {

	var listInput todo.TodoList
	if err := context.BindJSON(&listInput); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	userId := getUserId(context)
	listId, err := handler.service.ListManager.CreateList(listInput, userId)

	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

type GetAllListsResponse struct {
	Lists []todo.TodoList `json:"data"`
}

func (handler *Handler) getAllLists(context *gin.Context) {
	userId := getUserId(context)

	lists, err := handler.service.ListManager.GetAllLists(userId)

	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, GetAllListsResponse{
		Lists: lists,
	})
}

func (handler *Handler) getListById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, "invalid id")
		return
	}
	list, err := handler.service.ListManager.GetListById(id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, list)
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (handler *Handler) updateList(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, "invalid id")
		return
	}
	var listInput UpdateListInput
	if err := context.BindJSON(&listInput); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	err = handler.service.ListManager.UpdateList(id, todo.UpdateListInput(listInput))

	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (handler *Handler) deleteList(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, "invalid id")
		return
	}
	err = handler.service.ListManager.DeleteList(id)

	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
