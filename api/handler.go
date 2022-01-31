package api

import (
	"github.com/eugeneuskov/log-service/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
	mode     string
}

func NewHandler(services *services.Services, mode string) *Handler {
	return &Handler{services, mode}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(h.mode)

	router := gin.New()

	router.GET("/info", h.info)

	return router
}
