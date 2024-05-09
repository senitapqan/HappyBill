package handler

import (
	"errors"
	"happyBill/dtos"
	"strconv"
	"time"

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

func ValidateSearch(c *gin.Context, search *dtos.Search) error {
	regionId, err := strconv.Atoi(c.Param("region"))

	if err != nil {
		return errors.New("Region id need to be int")
	}

	if regionId <= 0 {
		return errors.New("Region id cannot be negative")
	}

	checkIn, err := time.Parse("2006-01-02", c.Param("check_in"))

	if err != nil {
		return errors.New("Wrong format of check-in")
	}

	checkOut, err := time.Parse("2006-01-02", c.Param("check_out"))

	if err != nil {
		return errors.New("Wrong format of check-out")
	}

	search.CheckIn = checkIn.String()
	search.CheckOut = checkOut.String()
	search.RegionId = regionId

	return nil

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

func ValidateLike(c *gin.Context) (string, error) {
	likeAction := c.Query("action")
	if likeAction == "like" || likeAction == "unlike" {
		return likeAction, nil
	}
	return "", errors.New("wrong format of action")
}

func ValidateStatus(c *gin.Context) (string, error) {
	status := c.DefaultQuery("status", "active")
	if status == "active" || status == "pending" || status == "passed" {
		return status, nil
	}
	return "", errors.New("wrong format of action")
}
