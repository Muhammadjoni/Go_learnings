package controllers

import (
	"net/http"
	"strings"

	"database/sql"
	"fmt"
	models "gin-todo-prc/models"
	"log"

	"github.com/gin-gonic/gin"

	//"gorm.io/gorm"

	_ "github.com/lib/pq"
)

// database assignation
var db *sql.DB

func CreateATodo(c *gin.Context) {
	item := c.Param("task")

	//Item validation
	if len(item) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"mess": "Enter an item"})
	} else {
		//Create todo
		var TodoItem models.Todo

		TodoItem.Task = item
		TodoItem.Done = false

		//insert to db
		_, err := db.Query("INSERT INTO list(item, done) VALUES($1, $2);", TodoItem.Task, TodoItem.Done)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error in database"})
		}

		//log msg
		log.Println("created todo item", item)
		//return resp
		c.JSON(http.StatusCreated, gin.H{"items": &TodoItem})
	}
}

// Get all the Todos
func GetAllTodos(c *gin.Context) {
	//Get all the rows using SelectQuery
	rows, err := db.Query("Select & From list")
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "error with db"})
	}

	// items << rows
	items := make([]models.Todo, 0)

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			item := models.Todo{}
			if err := rows.Scan(&item.Id, &item.Task, &item.Done); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "error with db"})
			}
			item.Task = strings.TrimSpace(item.Task)
			items = append(items, item)
		}
	}
	// return response
	c.JSON(http.StatusOK, gin.H{"items": items})

}

// Edit the Todo by Id
func UpdateTodoById(c *gin.Context) {
	id := c.Param("id")
	done := c.Param("done")

	// Validate id and done
	if len(id) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"msg": "id, please?"})
	} else if len(done) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"msg": "please enter a status"})
	} else {

		// Find and update the todo item
		var exists bool
		err := db.QueryRow("SELECT * FROM list WHERE id=$1;", id).Scan(&exists)
		if err != nil && err == sql.ErrNoRows {
			fmt.Println(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"msg": "not found"})
		} else {
			_, err := db.Query("UPDATE list SET done=$1 WHERE id=$2;", done, id)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "error with DB"})
			}

			// Log message
			log.Println("updated todo item", id, done)

			// Return success response
			c.JSON(http.StatusOK, gin.H{"msg": "successfully updated todo item", "todo": id})
		}
	}
}

// Delete the Todo by Id
func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")

	// Validate id
	if len(id) == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "please enter an id"})
	} else {
		// Find and delete the todo item
		var exists bool
		err := db.QueryRow("SELECT * FROM list WHERE id=$1;", id).Scan(&exists)
		if err != nil && err == sql.ErrNoRows {
			fmt.Println(err.Error())
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			_, err = db.Query("DELETE FROM list WHERE id=$1;", id)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			}

			// Log message
			log.Println("deleted todo item", id)

			// Return success response
			c.JSON(http.StatusOK, gin.H{"message": "successfully deleted todo item", "todo": id})
		}
	}
}

// // Get the Todo by Id
// func GetATodoById(c *gin.Context) {

// }
