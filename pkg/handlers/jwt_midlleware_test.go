package handlers

import (
	"errors"
	"github.com/ainurqa95/todo-list-app/pkg/services"

	mock_service "github.com/ainurqa95/todo-list-app/pkg/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestJwt_UserIdentity(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, token string)

	tests := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `1`,
		},
		{
			name:                 "Empty header",
			headerName:           "Authorization",
			headerValue:          "",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "Wrong bearer",
			headerName:           "Authorization",
			headerValue:          "Bersre token",
			token:                "token",
			mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:        "Wrong token",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
				r.EXPECT().ParseToken(token).Return(1, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mock_service.NewMockAuthorization(controller)
			test.mockBehavior(repo, test.token)

			services := &services.Service{Authorization: repo}
			handler := NewHandler(services)

			// Init Endpoint
			r := gin.New()
			r.GET("/identity", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, "%d", id)
			})

			// Init Test Request
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(recorder, req)

			// Asserts
			assert.Equal(t, recorder.Code, test.expectedStatusCode)
			assert.Equal(t, recorder.Body.String(), test.expectedResponseBody)
		})
	}
}
func TestHandler_getUser(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(userCtx, id)
		return ctx
	}

	testTable := []struct {
		name       string
		ctx        *gin.Context
		id         int
		shouldFail bool
	}{
		{
			name: "Ok",
			ctx:  getContext(1),
			id:   1,
		},
		{
			name: "Empty",
			ctx:  &gin.Context{},
			id:   0,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			id := getUserId(test.ctx)
			assert.Equal(t, id, test.id)
		})
	}
}
