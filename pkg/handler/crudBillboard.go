package handler

import (
	"happyBill/dtos"
	"happyBill/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// @Summary		Create
// @Tags			admin/billboard
// @Security		apiKeyAuth
// @Description	Create the billboard and add it to data base
// @ID				create-billboard
// @Accept			json
// @Produce		json
// @Param			input	body	models.Product	true	" height / width / display_type / location_id / price"
// @Router			/admin/bill [post]
func (h *Handler) createBillboard(c *gin.Context) {
	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Info().Msg("started handling create billboard request")
	id, err := h.service.CreateBillboard(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "New billboard was created!",
		"id":      id,
	})
}

// @Summary		GetAll
// @Security		ApiKeyAuth
// @Tags			admin/billboard
// @Description	Get all billboards from data base
// @ID				get-billboards
// @Accept			json
// @Produce		json
// @Router			/admin/bill [get]
func (h *Handler) getAllBillboards(c *gin.Context) {
	page, err := h.validator.ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	clinetId, err := getId(c, clientCtx)
	if err != nil {
		clinetId = -1
	}
	log.Info().Msg("started handling get all billboards request")

	products, pagination, err := h.service.GetAllBillboards(page, clinetId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetAllBillboardsResponse{
		Data:       products,
		Pagination: pagination,
	})
}

func (h *Handler) getAllSearchedBillboards(c *gin.Context) {
	page, err := h.validator.ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var search dtos.Search
	if err := h.validator.ValidateSearch(c, &search); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var filter dtos.Filter
	if err := h.validator.ValidateFilter(c, &filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling get all billboards request")

	products, pagination, err := h.service.GetAllSearchedBillboards(page, search, filter)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetAllBillboardsResponse{
		Data:       products,
		Pagination: pagination,
	})
}

// @Summary		GetById
// @Tags			admin/billboard
// @Security		ApiKeyAuth
// @Description	Get the billboard from data base
// @ID				get-billboard
// @Accept			json
// @Produce		json
// @Router			/admin/bill/{id} [get]
func (h *Handler) getBillboardById(c *gin.Context) {
	id, err := h.validator.ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	log.Info().Msg("started handling get billboard by id request")
	product, err := h.service.GetBillboardById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetBillboardByIdResponse{
		Data: product,
	})
}

func (h *Handler) getMyBillboards(c *gin.Context) {
	clientId, _ := getId(c, clientCtx)

	page, err := h.validator.ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling get my billboards request")
	products, pagination, err := h.service.GetMyBillboards(clientId, page)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dtos.GetAllBillboardsResponse{
		Data:       products,
		Pagination: pagination,
	})
}

func (h *Handler) likeBillboard(c *gin.Context) {
	clientId, _ := getId(c, clientCtx)
	productId, err := h.validator.ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	action, err := h.validator.ValidateLike(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling like some billboard request")
	err = h.service.LikeBillboard(clientId, productId, action)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "query was successfully pended",
	})
}

// @Summary		UpdateById
// @Tags			admin/billboard
// @Security		ApiKeyAuth
// @Description	Update
// @ID				update-billboard
// @Accept			json
// @Produce		json
// @Router			/admin/bill/{id} [put]
func (h *Handler) updateBillboard(c *gin.Context) {
	id, err := h.validator.ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter: "+err.Error())
		return
	}

	var input dtos.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling update billboard request")

	if err := h.service.UpdateBillboard(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Updated succesfully",
	})

}

func (h *Handler) deleteBillboard(c *gin.Context) {
	id, err := h.validator.ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling delete billboard request")

	err = h.service.DeleteBillboard(id)

	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
