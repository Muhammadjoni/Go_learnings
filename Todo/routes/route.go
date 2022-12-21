package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	url := r.Group("/url")
	{
		url.GET("todo")
		url.POST("todo")
		url.GET("todo/:id")
		url.PUT("todo/:id")
		url.DELETE("todo/:id")
	}
	return r
}
