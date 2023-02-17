package main

import (
	"auth_prtc/cmd"
	"auth_prtc/pkg/infra/util"
	"log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	db *sqlx.DB
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

	//2
	cmd.Execute()

}
