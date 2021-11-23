package handlers

import (
	"github.com/ainurqa95/todo-list-app/pkg/services"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/ainurqa95/todo-list-app/docs"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	api := router.Group("/api", handler.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", handler.createList)
			lists.GET("/", handler.getAllLists)
			lists.GET("/:id", handler.getListById)
			lists.PUT("/:id", handler.updateList)
			lists.DELETE("/:id", handler.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", handler.createItem)
				items.GET("/", handler.getAllItems)
				// items.GET("/:item_id", h.getItemById)
				// items.PUT("/:item_id", h.updateItem)
				// items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
