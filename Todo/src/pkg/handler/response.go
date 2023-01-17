package handler

import (
	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	// zap.Error(message)
	zap.String("error", message)

	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
