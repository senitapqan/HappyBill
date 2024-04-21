package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
		log.Info().Msg(fmt.Sprintf("Here is header: %s", header))
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

		log.Info().Msg("PARSING TOKEN")

		userId, roles, err := h.service.ParseToken(headerParts[1])

		log.Info().Msg(fmt.Sprintf("userId: %d", userId))

		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Request.Header.Add(userCtx, strconv.Itoa(userId))
		for _, role := range roles {
			c.Request.Header.Add(role.Role, strconv.Itoa(role.Id))
		}

		c.Next()
	}
}

func (h *Handler) adminIdentify() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getId(c, adminCtx)	
		if err != nil {
			log.Error().Msg("You dont have admin permission")
			newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
			return
		}

		log.Info().Msg(fmt.Sprintf("ADMIN WITH ID %d SENT REQUEST", id))
		c.Next()
	}
}

func (h *Handler) clientIdentify() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getId(c, clientCtx)	
		if err != nil {
			log.Error().Msg("You are not signed in")
			newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
			return
		}

		log.Info().Msg(fmt.Sprintf("CLIENT WITH ID %d SENT REQUEST", id))
		c.Next()
	}
}

func getId(c *gin.Context, header string) (int, error) {
	log.Info().Msg("GETTING HEADER OF THE REQUEST")

	id := c.GetHeader(header)
	if id == "" {
		return 0, fmt.Errorf("%s id not found", header)
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return 0, fmt.Errorf("cant converse %s id", header)
	}

	return intId, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
