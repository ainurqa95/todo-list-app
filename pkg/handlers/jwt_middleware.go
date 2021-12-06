package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	if headerParts[0] != "Bearer" || len(headerParts) != 2 {
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
		return 0
	}
	return userId.(int)
}
