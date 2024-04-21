package handler

import (
	"fmt"
	"happyBill/dtos"
	"happyBill/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//	@Summary		Create
//	@Tags			admin/billboard
//	@Security		apiKeyAuth
//	@Description	Create the billboard and add it to data base
//	@ID				create-billboard
//	@Accept			json
//	@Produce		json
//	@Param			input	body	models.Product	true	" height / width / display_type / location_id / price"
//	@Router			/admin/bill [post]
func (h *Handler) createBillboard(c *gin.Context) {
	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		log.Error().Msg("STARTED HANDLING CREATE BILLBOARD REQUEST")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Info().Msg("STARTED HANDLING CREATE BILLBOARD REQUEST")
	id, err := h.service.CreateBillboard(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "New billboard was created!",
		"id":      id,
	})

	log.Info().Msg("CREATE BILLBOARD REQUEST ENDED")

}

//	@Summary		GetAll
//	@Security		ApiKeyAuth
//	@Tags			admin/billboard
//	@Description	Get all billboards from data base
//	@ID				get-billboards
//	@Accept			json
//	@Produce		json
//	@Router			/admin/bill [get]
//host:port/admin/bill?page=1&limit=10&q="naruto"
func (h *Handler) getAllBillboards(c *gin.Context) {
	log.Info().Msg("STARTED HANDLING GET ALL BILLBOARDS REQUEST")

	page, err := ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.service.GetAllBillboards(page)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetAllBillboardsResponse{
		Data: products,
	})

	log.Info().Msg("GET ALL BILLBOARDS REQUEST ENDED")
}

//	@Summary		GetById
//	@Tags			admin/billboard
//	@Security		ApiKeyAuth
//	@Description	Get the billboard from data base
//	@ID				get-billboard
//	@Accept			json
//	@Produce		json
//	@Router			/admin/bill/{id} [get]
func (h *Handler) getBillboardById(c *gin.Context) {
	id, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	log.Info().Msg("STARTED HANDLING GET BILLBOARD BY ID REQUEST")

	product, err := h.service.GetBillboardById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetBillboardByIdResponse{
		Data: product,
	})

	log.Info().Msg("GET BILLBOARD BY ID REQUEST ENDED")
}

func (h *Handler) getMyBillboards(c *gin.Context) {
	clientId, _ := getId(c, clientCtx)

	page, err := ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.service.GetMyBillboards(clientId, page)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dtos.GetAllBillboardsResponse{
		Data: products,
	})
}

//	@Summary		UpdateById
//	@Tags			admin/billboard
//	@Security		ApiKeyAuth
//	@Description	Update
//	@ID				update-billboard
//	@Accept			json
//	@Produce		json
//	@Router			/admin/bill/{id} [put]
func (h *Handler) updateBillboard(c *gin.Context) {
	id, err := ValidateId(c)

	if err != nil {

		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input models.Product

	log.Info().Msg(fmt.Sprintf("input.Height: + %d", input.Height))

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("STARTED HANDLING UPDATE BILLBOARD REQUEST")

	if err := h.service.UpdateBillboard(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Updated succesfully",
	})

	log.Info().Msg("UPDATE BILLBOARD REQUEST ENDED")

}

func (h *Handler) deleteBillboard(c *gin.Context) {
	id, err := ValidateId(c)

	if err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("STARTED HANDLING DELETE BILLBOARD REQUEST")

	err = h.service.DeleteBillboard(id)

	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

	log.Info().Msg("DELETE BILLBOARD REQUEST ENDED")

}
