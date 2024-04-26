package handler

import (
	"happyBill/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) getMyProfile(c *gin.Context) {
	userId, _ := getId(c, userCtx)
	log.Info().Msg("started handling get my profile request")
	user, err := h.service.GetClientByUserId(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetUserByIdResponse{
		Data: user,
	})
}

func (h *Handler) GetClientById(c *gin.Context) {
	clientId, err := ValidateId(c)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling get client by Id request")
	client, err := h.service.GetClientById(clientId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetUserByIdResponse{
		Data: client,
	})
}

func (h *Handler) updateMyProfile(c *gin.Context) {
	userId, _ := getId(c, userCtx)

	var input dtos.UpdateUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("started handling update my profile request")
	err := h.service.UpdateMyProfile(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Updated succesfully",
	})
}
