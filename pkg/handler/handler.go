package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maXvostik/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	auth := routes.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := routes.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updeteList)
			lists.DELETE("/:id", h.delitList)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updeteItem)
				items.DELETE("/:item_id", h.delitItem)
			}
		}
	}
	return routes
}
