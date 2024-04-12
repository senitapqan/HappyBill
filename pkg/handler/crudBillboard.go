package handler

import (
	"happyBill/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllBillboardsResponse struct {
	Data []models.Product `json:"data"`
}

type getBillboardByIdResponse struct {
	Data models.Product `json:"data"`
}

func (h *Handler) createBillboard(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateBillboard(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message: ": "New billboard was created!",
		"id":        id,
	})

}

func (h *Handler) getAllBillboards(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.service.GetAllBillboards()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBillboardsResponse{
		Data: products,
	})

}

func (h *Handler) getBillboardById(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	product, err := h.service.GetBillboardById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getBillboardByIdResponse{
		Data: product,
	})

}

func (h *Handler) updateBillboard(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input models.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateBillboard(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteBillboard(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteBillboard(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
