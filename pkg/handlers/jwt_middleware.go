package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (handler *Handler) userIdentity(context *gin.Context) {
	header := context.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(context, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(context, http.StatusUnauthorized, headerParts[0])
		newErrorResponse(context, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := handler.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	context.Set(userCtx, userId)
}

func getUserId(context *gin.Context) int {
	userId, ok := context.Get(userCtx)
	if !ok {
		newErrorResponse(context, http.StatusUnauthorized, "user not found")
		return 0
	}
	return userId.(int)
}
