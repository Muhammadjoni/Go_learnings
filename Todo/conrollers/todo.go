package Conrollers

import (
	"net/http"

	Models "gin-todo-prc/models"

	"github.com/gin-gonic/gin"
)

//Create

func CreateATodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)
	err := Models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
