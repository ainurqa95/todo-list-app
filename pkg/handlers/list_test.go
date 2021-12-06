package handlers

import (
	"bytes"
	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/services"
	mock_service "github.com/ainurqa95/todo-list-app/pkg/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateList(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockListManager, list todo.TodoList, userId int)

	tests := []struct {
		name                 string
		inputBody            string
		inputList            todo.TodoList
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"title": "My title", "description": "My description"}`,
			inputList: todo.TodoList{
				Id: 1,
				Title:       "My title",
				Description: "My description",
			},
			mockBehavior: func(r *mock_service.MockListManager, list todo.TodoList, userId int) {},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mock_service.NewMockListManager(controller)
			test.mockBehavior(repo, test.inputList, 1)

			services := &services.Service{ListManager: repo}
			handler := Handler{services}

			// Init Endpoint
			ginServer := gin.New()
			ginServer.POST("/api/lists/", handler.createList)

			// Create Request
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/lists/",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			ginServer.ServeHTTP(recorder, req)

			// Assert
			//assert.Equal(t, recorder.Code, test.expectedStatusCode)
			//assert.Equal(t, recorder.Body.String(), test.expectedResponseBody)
		})
	}
}
