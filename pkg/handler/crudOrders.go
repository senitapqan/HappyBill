package handler

import (
	"happyBill/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createOrder(c *gin.Context) {

}

func (h *Handler) getAllOrders(c *gin.Context) {
	page, err := ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var orders []dtos.Order

	orders, err = h.service.GetAllOrders(page)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetOrdersResponse{
		Data: orders,
	})
}

//	@Summary		Get My Orders
//	@Security		ApiKeyAuth
//	@Tags			profile/my-orders
//	@Description	Get all orders i have 
//	@ID				get-my-orders
//	@Accept			json
//	@Produce		json
//	@Params			page query string int "The page in which now I am"
//	@Router			/profile/my-orders/ [get]
func (h *Handler) getMyOrders(c *gin.Context) {
	clientId, _ := getId(c, clientCtx)
	page, err := ValidatePage(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var myOrders []dtos.MyOrder
	myOrders, err = h.service.GetMyOrders(clientId, page)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetMyOrdersResponse{
		Data: myOrders,
	})
}
