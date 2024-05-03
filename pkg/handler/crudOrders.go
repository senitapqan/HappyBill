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
		newErrorResponse(c, http.StatusBadRequest, "invalid input body: " + err.Error())
		return
	}

	productId, err := ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id:" + err.Error())
		return
	}

	input.ProductId = productId

	log.Info().Msg("started handling create my order request")
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
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := ValidateStatus(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var myOrders []dtos.MyOrder
	log.Info().Msg("started handling get all my orders request")
	myOrders, err = h.service.GetMyOrders(clientId, page, status)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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

	log.Info().Msg("started handling get all orders request")
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


// @Summary		Get All Manager Orders
// @Security		ApiKeyAuth
// @Tags			manager/
// @Description	Get all manager orders from data base
// @ID				get-manager-orders
// @Accept			json
// @Produce		json
// @Router			/manager/ [get]
func (h *Handler) getAllManagerOrders(c *gin.Context) {
	page, err := ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Info().Msg("getting id of a manager")
	manager_id, err := getId(c, managerCtx)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	log.Info().Msg("started handling get all manager orders request")
	orders, err := h.service.GetAllManagerOrders(manager_id, page)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	log.Info().Msg("get all manager orders works properly")

	c.JSON(http.StatusOK, dtos.GetAllManagerOrdersResponse{
		Data: orders,
	})
}

// @Summary			Get Manager Order By Id
// @Security		ApiKeyAuth
// @Tags			manager
// @Description		Get the manager order from data base with ID
// @ID				get-manager-order
// @Accept			json
// @Produce			json
// @Router			/manager/:id [get]
func (h *Handler) getManagerOrderById(c *gin.Context) {
	id, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter: "+err.Error())
		return
	}

	log.Info().Msg("started handling get manager order by id request")
	manager, err := h.service.GetManagerOrderById(id)

	if err != nil {

		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetManagerOrderResponse{
		Data: manager,
	})
}

// @Summary			UpdateManagerOrder
// @Tags			manager/
// @Security		ApiKeyAuth
// @Description		Update
// @ID				update-manager-order
// @Accept			json
// @Produce			json
// @Router			/manager/{id} [put]
func (h *Handler) updateManagerOrder(c *gin.Context) {
	id, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter: "+err.Error())
		return
	}

	var input dtos.UpdateOrder
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling update manager order request")

	if err := h.service.UpdateManagerOrder(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Updated succesfully",
	})

}

