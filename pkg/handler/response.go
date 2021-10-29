package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:Message`
}

func newErrorMessage(context *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	context.AbortWithStatusJSON(statusCode, Error{Message: message})
}
