package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	zap.String("error", message)

	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
