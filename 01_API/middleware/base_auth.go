package middleware

import "github.com/gin-gonic/gin"

func BaseAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"whoCares": "ShouldBeToughPassword",
	})
}
