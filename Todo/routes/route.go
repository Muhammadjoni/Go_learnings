package routes

import (
	controllers "gin-todo-prc/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	url := r.Group("/url")
	{
		url.GET("/todos", controllers.GetAllTodos)
		url.POST("/todo", controllers.CreateATodo)
		// url.GET("todo/:id", controllers.GetATodoById)
		url.PUT("todo/:id", controllers.UpdateTodoById)
		url.DELETE("todo/:id", controllers.DeleteTodoById)
	}

	return r
}
