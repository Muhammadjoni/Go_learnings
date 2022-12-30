package service

import (
	"context"

	"gin-todo-prc/dtb"

	"github.com/BetaLixT/tsqlx"
)

// func InitializeApp() (*app, error) {
// 	// loggerFactory, err := logger.NewLoggerFactory()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	stt := NewApp()
// 	return stt, nil
// }

// func Start() {
// // function to start the service that runs migration
//     app, err := util.SetupPostgres()
// 	app.startService()
// }

type app struct {
	dbctx *tsqlx.TracedDB
}

// func NewApp(
// 	dbctx *tsqlx.TracedDB,
// ) *app {
// 	return &app{
// 		dbctx: dbctx,
// 	}
// }

func (a *app) startService() {
	// cx := context.TODO()
	dtb.RunMigrations(
		context.TODO(),
		a.dbctx,
		dtb.GetMigrationScripts(),
	)

}
