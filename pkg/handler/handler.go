package handler

import (
	"happyBill/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func (h Handler) InitRoutes() *gin.Engine{
	router := gin.New()
	
	return router
}

func NewHandler(serv service.Service) *Handler {
	return &Handler{
		service: serv,
	}
}
