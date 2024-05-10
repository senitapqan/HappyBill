package handler

import (
	"happyBill/dtos"
	"happyBill/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// @Summary		Create Manager
// @Security		ApiKeyAuth
// @Tags			admin/manager
// @Description	Create new manager to Data Base
// @ID				create-manager
// @Accept			json
// @Produce		json
// @Router			/admin/admin [post]
func (h *Handler) createManager(c *gin.Context) {
	var request models.User
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Info().Msg("started handling create manager request")
	id, err := h.service.CreateManager(request)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"new Manager was succesfully added with id": id,
	})

}

// @Summary		Get all Managers
// @Security		ApiKeyAuth
// @Tags			admin/manager
// @Description	Get all managers from data base
// @ID				get-managers
// @Accept			json
// @Produce		json
// @Router			/admin/manager [get]
func (h *Handler) getAllManager(c *gin.Context) {
	page, err := h.validator.ValidatePage(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling get all managers request")
	managers, pagination, err := h.service.GetAllManagers(page)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	log.Info().Msg("get all managers works properly")

	c.JSON(http.StatusOK, dtos.GetAllManagersResponse{
		Data: managers,
		Pagination: pagination,
	})
}

// @Summary		Get Manager By Id
// @Security		ApiKeyAuth
// @Tags			admin/manager
// @Description	Get the manager from data base with ID
// @ID				get-manager
// @Accept			json
// @Produce		json
// @Router			/admin/manager/:id [get]
func (h *Handler) getManagerById(c *gin.Context) {
	id, err := h.validator.ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter: "+err.Error())
		return
	}

	log.Info().Msg("started handling get manager by id request")
	manager, err := h.service.GetManagerById(id)

	if err != nil {

		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetManagerResponse{
		Data: manager,
	})
}

func (h *Handler) deleteManager(c *gin.Context) {

}

func (h *Handler) updateManager(c *gin.Context) {

}
