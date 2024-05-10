package handler

import (
	"errors"
	"happyBill/dtos"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=validator.go -destination=validatorMock/mock.go
type Validator interface {
	ValidateSearch(c *gin.Context, search *dtos.Search) error
	ValidateFilter(c *gin.Context, filter *dtos.Filter) error 
	ValidateId(c *gin.Context) (int, error) 
	ValidatePage(c *gin.Context) (int, error) 
	ValidateLike(c *gin.Context) (string, error)
	ValidateStatus(c *gin.Context) (string, error)
}

type validator struct {

}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) ValidateSearch(c *gin.Context, search *dtos.Search) error {
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

func (v *validator) ValidateFilter(c *gin.Context, filter *dtos.Filter) error {

	return nil
}

func (v *validator) ValidateId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1, err
	}

	if id <= 0 {
		return -1, errors.New("id cannot be negative")
	}
	return id, nil
}

func (v *validator) ValidatePage(c *gin.Context) (int, error) {
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

func (v *validator) ValidateLike(c *gin.Context) (string, error) {
	likeAction := c.Query("action")
	if likeAction == "like" || likeAction == "unlike" {
		return likeAction, nil
	}
	return "", errors.New("wrong format of action")
}

func (v *validator) ValidateStatus(c *gin.Context) (string, error) {
	status := c.DefaultQuery("status", "active")
	if status == "active" || status == "pending" || status == "passed" {
		return status, nil
	}
	return "", errors.New("wrong format of action")
}
