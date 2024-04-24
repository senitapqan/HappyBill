package handler

import (
	"fmt"
	"happyBill/dtos"
	"happyBill/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) createMyOrder(c *gin.Context) {
	clientId, _ := getId(c, clientCtx)

	var input models.Order
	if err := c.BindJSON(&input); err != nil {
		log.Error().Msg("Something wrong with request body params: " + err.Error())
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	productId, err := ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id:" + err.Error())
		return
	}

	input.ProductId = productId

	id, err := h.service.CreateOrder(clientId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Order with id %d was created and sent to manager for pending", id),
	})
}

func (h *Handler) deleteMyOrder(c *gin.Context) {

}

// @Summary		Get My Orders
// @Security		ApiKeyAuth
// @Tags			profile/my-orders
// @Description	Get all orders i have
// @ID				get-my-orders
// @Accept			json
// @Produce		json
// @Params			page query string int "The page in which now I am"
// @Router			/profile/my-orders/ [get]
func (h *Handler) getMyOrders(c *gin.Context) {
	clientId, _ := getId(c, clientCtx)
	page, err := ValidatePage(c)
	if err != nil {
		log.Error().Msg("Errors with page or client?")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := ValidateStatus(c)
	if err != nil {
		log.Error().Msg("Status of orders is inccorrect")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Info().Msg(fmt.Sprintf("Client with clientId %d want to see their orders on %d page", clientId, page))

	var myOrders []dtos.MyOrder
	myOrders, err = h.service.GetMyOrders(clientId, page, status)

	if err != nil {
		log.Error().Msg("Seems like in service were some errors")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg("Everything is good")
	c.JSON(http.StatusOK, dtos.GetMyOrdersResponse{
		Data: myOrders,
	})
}

// @Summary		Get All Orders
// @Security		ApiKeyAuth
// @Tags			admin/order
// @Description	Get all orders from data base
// @ID				get-all-orders
// @Accept			json
// @Produce		json
// @Router			/admin/order/ [get]
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

func (h *Handler) deleteOrder(c *gin.Context) {

}
