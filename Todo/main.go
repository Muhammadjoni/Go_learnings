package main

import (
	"gin-todo-prc/cmd"
	"gin-todo-prc/routes"

	_ "github.com/lib/pq"
)

func main() {
	// SetupPostgres()
	cmd.Execute()
	r := routes.SetupRouter()
	r.Run(":8080")
}
