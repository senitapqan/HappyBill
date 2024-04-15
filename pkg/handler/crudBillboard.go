package handler

import (
	"fmt"
	"happyBill/dtos"
	"happyBill/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

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

	log.Info().Msg(fmt.Sprintf("STARTED HANDLING CREATE BILLBOARD REQUEST"))
	id, err := h.service.CreateBillboard(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message: ": "New billboard was created!",
		"id":        id,
	})

	log.Info().Msg(fmt.Sprintf("CREATE BILLBOARD REQUEST ENDED"))

}

func (h *Handler) getAllBillboards(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("STARTED HANDLING GET ALL BILLBOARDS REQUEST"))

	products, err := h.service.GetAllBillboards()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetAllBillboardsResponse{
		Data: products,
	})

	log.Info().Msg(fmt.Sprintf("GET ALL BILLBOARDS REQUEST ENDED"))

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

	log.Info().Msg(fmt.Sprintf("STARTED HANDLING GET BILLBOARD BY ID REQUEST"))

	product, err := h.service.GetBillboardById(id)

	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetBillboardByIdResponse{
		Data: product,
	})

	log.Info().Msg(fmt.Sprintf("GET BILLBOARD BY ID REQUEST ENDED"))

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

	log.Info().Msg(fmt.Sprintf("STARTED HANDLING UPDATE BILLBOARD REQUEST"))

	if err := h.service.UpdateBillboard(id, input); err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Updated succesfully",
	})

	log.Info().Msg(fmt.Sprintf("UPDATE BILLBOARD REQUEST ENDED"))

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

	log.Info().Msg(fmt.Sprintf("STARTED HANDLING DELETE BILLBOARD REQUEST"))

	err = h.service.DeleteBillboard(id)

	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

	log.Info().Msg(fmt.Sprintf("DELETE BILLBOARD REQUEST ENDED"))

}
