package handler

import (
	"happyBill/dtos"
	"happyBill/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// @Summary SignIn
// @Tags auth
// @Description login to account
// @ID login-account
// @Accept json
// @Produce json
// @Param input body dtos.SignInRequest true "username / password"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var request dtos.SignInRequest
	if err := c.BindJSON(&request); err != nil {
		log.Error().Msg("Error binding JSON")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info().Msg("JSON binded successfully")

	token, err := h.service.GenerateToken(request.Username, request.Password)

	if err != nil {
		log.Error().Msg("Error generating token")
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	log.Info().Msg("Token generated")
	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	var request models.User
	if err := c.BindJSON(&request); err != nil {
		log.Error().Msg("Error binding JSON")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	log.Info().Msg("JSON binded successfully")

	id, err := h.service.CreateClient(request)

	if err != nil {
		log.Error().Msg("Error crating a client")
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	log.Info().Msg("Client created successfully")

	c.JSON(http.StatusOK, map[string]int{
		"new user was succesfully added with id": id,
	})
}
