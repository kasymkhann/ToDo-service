package handler

import (
	ser "to-doProjectGo/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *ser.Service
}

func NewHandler(service *ser.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) ThisRouter() *gin.Engine {

	router := gin.New()

	entr := router.Group("/entr")
	{
		entr.POST("/sign-up", h.SignUp)
		entr.POST("/sign-in", h.SignIn)
	}
	api := entr.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createLists)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListsId)
			lists.PUT("/:id", h.putLists)
			lists.DELETE("/:id", h.deleteLists)

			items := lists.Group("/:id")
			{
				items.POST("/", h.createItems)
				items.GET("/", h.getAllItems)

			}
		}
		items := api.Group("items")
		{
			items.GET("/:id", h.getItemsId)
			items.PUT("/:id", h.putItems)
			items.DELETE("/:id", h.deleteItems)
		}

	}
	return router
}
