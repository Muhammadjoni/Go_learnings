package main

import (
	"golang-gin-prc/controller"
	"golang-gin-prc/middleware"
	"golang-gin-prc/service"
	"io"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create(("gin.log")) // creatig the dfile using OS library
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger())

	// server.GET("/testingApi", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OKKKK!!!",
	// 	})
	// })

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run("")
}
