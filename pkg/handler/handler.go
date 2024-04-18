 package handler

import (
	"happyBill/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	_ "happyBill/docs"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		admin.Use(h.adminIdentify())

		billboard := admin.Group("/bill")
		{
			billboard.GET("/", h.getAllBillboards)
			billboard.GET("/:id", h.getBillboardById)
			billboard.POST("/", h.createBillboard)
			billboard.DELETE("/:id", h.deleteBillboard)
			billboard.PUT("/:id", h.updateBillboard)
		}

		managers := admin.Group("/manager")
		{
			managers.POST("/", h.createManager)
			managers.GET("/", h.getAllManager)
			managers.GET("/:id", h.getManagerById)
			managers.DELETE("/:id", h.deleteManager)
			managers.PUT("/:id", h.updateManager)
		}
	}

	client := router.Group("/home")
	{
		client.GET("/my-orders")
	}

	return router
}


