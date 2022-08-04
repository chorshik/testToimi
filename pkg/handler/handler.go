package handler

import (
	"github.com/chorshik/testToimi/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	 services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		lists := api.Group("/ads")
		{
			lists.POST("/create", h.createAd)
			lists.GET("/", h.getAllAds)
			lists.GET("/:id", h.getAdById)
		}

	}

	return router
}
