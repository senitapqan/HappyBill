package handler

import (
	"happyBill/dtos"
	"happyBill/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getMyProfile(c *gin.Context) {
	userId, _ := getId(c, userCtx)

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

	client, err := h.service.GetClientById(clientId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetUserByIdResponse{
		Data: client,
	})
}

func (h *Handler) UpdateMyProfile(c *gin.Context) {
	userId, _ := getId(c, userCtx)

	var input models.User

	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.UpdateMyProfile(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"Message": "Updated succesfully",
	})
}
