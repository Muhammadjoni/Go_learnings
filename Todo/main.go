package main

import (
	"gin-todo-prc/cmd"
	"gin-todo-prc/src/pkg/handler"
	"gin-todo-prc/src/pkg/repo"
	"gin-todo-prc/src/pkg/service"
	"gin-todo-prc/src/util"

	"log"

	// "gin-todo-prc/pkg/handler"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// type Server struct {
// 	httpServer *httpServer
// }

// func db() {
// 	// var (
// 	// 	db *sqlx.DB
// 	// 	err error
// 	// )

// 	// util.SetupDb()
// 	// repo := repo.NewRepo(db)
// 	// services := service.NewService(repo)
// 	// handler := handler.NewHandler(services)

// }
var (
	db  *sqlx.DB
	err error
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cant init zap logger: %v", err)
	}

	defer logger.Sync()
	//1
	if db, err = util.ConnectToDb(); err != nil {
		zap.Error(err)
	}

	repo := repo.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	//2
	cmd.Execute()

	r := handler.SetupRouter()
	r.Run(":8080")
	// r := routes.SetupRouter()
}
