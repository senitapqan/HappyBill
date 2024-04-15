package handler

import (
	"happyBill/dtos"
	"happyBill/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) createManager(c *gin.Context) {
	/*_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	*/
	var request models.User
	if err := c.BindJSON(&request); err != nil {
		log.Error().Msg("Error binding JSON")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Info().Msg("JSON binded successfully")

	id, err := h.service.CreateManager(request)

	if err != nil {
		log.Error().Msg("Error creating manager")
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong: "+err.Error())
		return
	}

	log.Info().Msg("Manager created successfully")

	c.JSON(http.StatusOK, map[string]int{
		"new Manager was succesfully added with id": id,
	})

}

func (h *Handler) getAllManager(c *gin.Context) {
	/*_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	*/

	managers, err := h.service.GetAllManagers()

	if err != nil {
		log.Error().Msg("Error getting all managers")
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	log.Info().Msg("get all managers works properly")

	c.JSON(http.StatusOK, dtos.GetAllManagersResponse{
		Data: managers,
	})
}

func (h *Handler) getManagerById(c *gin.Context) {
	/*_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	*/
	id, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	log.Info().Msg("STARTED HANDLING GET MANAGER BY ID REQUEST")

	manager, err := h.service.GetManagerById(id)

	if err != nil {

		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetManagersResponse{
		Data: manager,
	})

	log.Info().Msg("GET MANAGER BY ID REQUEST ENDED")

}

func (h *Handler) deleteManager(c *gin.Context) {
	/*_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	*/
}

func (h *Handler) updateManager(c *gin.Context) {
	/*_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	*/
}
