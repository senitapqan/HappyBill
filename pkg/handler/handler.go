package handler

import (
	"happyBill/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	admin := router.Group("/admin")
	{
		admin.Use(h.userIdentify())
		admin.Use(h.roleIdentify(adminCtx))

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

		orders := admin.Group("/order")
		{
			orders.GET("/", h.getAllOrders)
		}
	}

	client := router.Group("")
	{
		profile := client.Group("/my")
		{
			profile.Use(h.userIdentify())
			profile.Use(h.roleIdentify(clientCtx))

			profile.GET("/", h.getMyProfile)
			profile.PUT("/", h.updateMyProfile)

			profile.GET("/my-orders", h.getMyOrders)
			profile.GET("/my-fav", h.getMyBillboards)

			profile.POST("/my-fav/:id", h.likeBillboard)

			profile.POST("/order/:id", h.createMyOrder)
			profile.DELETE("/order/:id", h.deleteMyOrder)

		}
		app := client.Group("/home")
		{
			app.GET("/", h.getAllBillboards)
			app.GET("/:id", h.getBillboardById)

		}
	}

	manager := router.Group("")
	{
		admin.Use(h.userIdentify())
		admin.Use(h.roleIdentify(managerCtx))

		manager.GET("/", h.getAllManagerOrders)
		manager.GET("/:id", h.getManagerOrderById)
		//manager.PUT("/:id", h.updateManagerOrder)
	}

	return router
}
