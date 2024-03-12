package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) createBillboard(c *gin.Context) {
	userId, roleId, err := h.getIds(adminCtx, c);
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) getBillboard(c *gin.Context) {
	userId, roleId, err := h.getIds(adminCtx, c);
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) updateBillboard(c *gin.Context) {
	userId, roleId, err := h.getIds(adminCtx, c);
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) deleteBillboard(c *gin.Context) {
	userId, roleId, err := h.getIds(adminCtx, c);
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}