package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	clientCtx           = "CLIENT"
	managerCtx          = "MANAGER"
	adminCtx            = "ADMIN"
)

func (h *Handler) userIdentify() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(c, http.StatusUnauthorized, "invalid header")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(c, http.StatusUnauthorized, "token is empty")
			return
		}

		userId, roles, err := h.service.ParseToken(headerParts[1])

		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, "Parsing works wrong: "+err.Error())
			return
		}

		c.Request.Header.Add(userCtx, strconv.Itoa(userId))
		for _, role := range roles {
			c.Request.Header.Add(role.Role, strconv.Itoa(role.Id))
		}

		c.Next()
	}
}
