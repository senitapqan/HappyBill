package handler

import (
	"happyBill/dtos"
	"happyBill/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) signIn(c *gin.Context) {
	var request dtos.SignInRequest
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Print("Sign in request")

	token, err := h.service.GenerateToken(request.Username, request.Password)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	var request models.User
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.CreateClient(request)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"new user was succesfully added with id": id,
	})
}
