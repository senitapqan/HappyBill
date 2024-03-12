package handler

import (
	"happyBill/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(serv service.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	admin := router.Group("/admin")
	{
		admin.Use(h.userIdentify())
		admin.GET("/getBill", h.getBillboard)
		admin.POST("/addBill", h.createBillboard)
		admin.DELETE("/deleteBill", h.deleteBillboard)
		admin.PUT("/updateBill", h.updateBillboard)
	}

	app := router.Group("/home")
	{
		app.Use(h.userIdentify())
	}

	return router
}

func (h *Handler) getIds(role string, c *gin.Context) (int, int, error) {
	userId, err := getId(c, userCtx)
	if err != nil {
		return -1, -1, err;
	}

	roleId, err := getId(c, role)
	
	return userId, roleId, err;
}
