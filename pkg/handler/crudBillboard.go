package handler

import (
	"fmt"
	"happyBill/dtos"
	"happyBill/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// @Summary Create
// @Tags admin/billboard
// @Description Create the billboard and add it to data base
// @ID create-billboard
// @Accept json
// @Produce json
// @Param input body models.Product true " height / width / display_type / location_id / price"
// @Router /admin/bill [post]
func (h *Handler) createBillboard(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		log.Error().Msg("Error getting admin id")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		log.Error().Msg("Error binding JSON")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("JSON binded successfully")

	id, err := h.service.CreateBillboard(input)
	if err != nil {
		log.Error().Msg("Error crating billboard")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg("billboard created successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message: ": "New billboard was created!",
		"id":        id,
	})

}

// @Summary GetAll
// @Tags admin/billboard
// @Description Get all billboards from data base
// @ID get-billboards
// @Accept json
// @Produce json
// @Router /admin/bill [get]
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

	c.JSON(http.StatusOK, dtos.GetAllBillboardsResponse{
		Data: products,
	})

}

// @Summary GetById
// @Tags admin/billboard
// @Description Get the billboard from data base
// @ID get-billboard
// @Accept json
// @Produce json
// @Router /admin/bill/{id} [get]
func (h *Handler) getBillboardById(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		log.Error().Msg("Error getting admin id")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ValidateId(c)

	if err != nil {
		log.Error().Msg("unvalid id")
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	product, err := h.service.GetBillboardById(id)

	if err != nil {
		log.Error().Msg("Error getting billboard by id")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg("get billboard by id works properly")

	c.JSON(http.StatusOK, dtos.GetBillboardByIdResponse{
		Data: product,
	})

}

// @Summary UpdateById
// @Tags admin/billboard
// @Description Update
// @ID update-billboard
// @Accept json
// @Produce json
// @Router /admin/bill/{id} [update]
func (h *Handler) updateBillboard(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		log.Error().Msg("Error getting admin id")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ValidateId(c)

	if err != nil {
		log.Error().Msg("unvalid id")
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input models.Product

	log.Info().Msg(fmt.Sprintf("input.Height: + %d", input.Height))

	if err := c.BindJSON(&input); err != nil {
		log.Error().Msg("Error binding JSON")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("JSON binded successfully")

	if err := h.service.UpdateBillboard(id, input); err != nil {
		log.Error().Msg("Error updating billboard")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg("billboard updated successfully")

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Updated succesfully",
	})

}

func (h *Handler) deleteBillboard(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		log.Error().Msg("Error getting admin id")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := ValidateId(c)

	if err != nil {
		log.Error().Msg("unvalid id")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteBillboard(id)

	if err != nil {
		log.Error().Msg("Error deleting billboard")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg("billboard deleted successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
