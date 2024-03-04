package handler

import (
	"happyBill/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	app := router.Group("/home") 
	{
		app.Use(h.userIdentify())
	}

	return router
}

func NewHandler(serv service.Service) *Handler {
	return &Handler{
		service: serv,
	}
}
