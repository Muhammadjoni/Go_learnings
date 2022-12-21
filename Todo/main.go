package main

import (
	"fmt"
	"gin-todo-prc/Config"
	Models "gin-todo-prc/models"
	"gin-todo-prc/routes"
	//"gorm.io/gorm"
)

var err error

func main() {

	//Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	//defer Config.DB.Close()

	Config.DB.AutoMigrate(Models.Todo{})

	r := routes.SetupRouter()

	r.Run()
}
