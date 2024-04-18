package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func ValidateId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1, err
	}

	if id <= 0 {
		return -1, errors.New("id cannot be negative")
	}
	return id, nil
}

func ValidatePage(c *gin.Context) (int, error) {
	pageStr := c.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(pageStr)

	if err != nil {
		return -1, errors.New("value of query parameter 'page' is not integer")
	}

	if pageInt <= 0 {
		return -1, errors.New("page cannot be negative")
	}

	return pageInt, nil
}
