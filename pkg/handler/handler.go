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

	unauth := router.Group("/")
	{
		unauth.GET("/home", h.getAllBillboards)
	}

	admin := router.Group("/admin")
	{
		
		admin.Use(h.userIdentify())
		bills := router.Group("/billboard")
		{
			bills.POST("/", h.createBillboard)
			bills.GET("/", h.getAllBillboards)
			bills.GET("/:id", h.getBillboardById)
			bills.PUT("/:id", h.updateBillboard)
			bills.DELETE("/:id", h.deleteBillboard)
		}
		admin.Use(h.userIdentify())

		admin.GET("/getBill", h.getAllBillboards)
		admin.GET("/getBill/:id", h.getBillboardById)
		admin.POST("/addBill", h.createBillboard)

		admin.DELETE("/deleteBill", h.deleteBillboard)
		admin.PUT("/updateBill", h.updateBillboard)
	}

	app := router.Group("/app")
	{
		app.Use(h.userIdentify())
		
		app.GET("")

	}

	return router
}

func (h *Handler) getIds(role string, c *gin.Context) (int, int, error) {
	userId, err := getId(c, userCtx)
	if err != nil {
		return -1, -1, err
	}

	roleId, err := getId(c, role)

	return userId, roleId, err
}
