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

	app := router.Group("/app")
	{
		app.Use(h.userIdentify())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
