package handlers

import (
	"net/http"
	"strconv"

	"github.com/ainurqa95/todo-list-app"
	"github.com/gin-gonic/gin"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description Create list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
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

// @Summary Get all lists
// @Security ApiKeyAuth
// @Tags lists
// @Description Get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} GetAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
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

// @Summary Get list by id
// @Security ApiKeyAuth
// @Tags lists
// @Description Get by list id
// @ID get-by-list-id
// @Accept  json
// @Produce  json
// @Param id path int true "List id"
// @Success 200 {object} todo.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [get]
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

// @Summary Update list by id
// @Security ApiKeyAuth
// @Tags lists
// @Description Updates list by id
// @ID update-list-id
// @Accept  json
// @Produce  json
// @Param id path int true "List id"
// @Param input body UpdateListInput true "update list info"
// @Success 200 {string} status "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [put]
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
