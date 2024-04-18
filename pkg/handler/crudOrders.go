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
		newErrorResponse(c, http.StatusBadGateway, err.Error() )
		return
	}

	c.JSON(http.StatusOK, dtos.GetAllOrdersResponse{
		Data: orders,
	})
}