package handler

import (
	"gin-todo-prc/docs"
	"gin-todo-prc/src/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		api := v1.Group("/api", h.userIdentity)
		{
			lists := api.Group("/lists")
			{
				lists.POST("", h.createList)
				lists.GET("", h.getAllList)
				lists.GET("/:id", h.getListByID)
				lists.PUT("/:id", h.updateList)
				lists.DELETE(":id", h.deleteList)

				items := lists.Group(":id/items")
				{
					items.POST("", h.createItem)
					items.GET("", h.getAllItem)
				}
			}
			items := api.Group("items")
			{
				items.GET("/:id", h.getItemByID)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}
	}

	return router
}
