package handler

import (
	"happyBill/dtos"
	"happyBill/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllMAnagersResponse struct {
	Data []dtos.User `json:"data"`
}

func (h *Handler) createManager(c *gin.Context) {
	/*_, _, err := h.getIds(adminCtx, c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	*/
	var request models.User
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.CreateManager(request)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong: "+err.Error())
		return
	}

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
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllMAnagersResponse{
		Data: managers,
	})

}
