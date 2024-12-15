package controllers

import (
	"attendance-app-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IShiftController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type ShiftController struct {
	service services.IShiftService
}

func NewShiftController(service services.IShiftService) IShiftController {
	return &ShiftController{service: service}
}

func (c *ShiftController) FindAll(ctx *gin.Context) {
	shifts, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shifts})
}

func (c *ShiftController) FindById(ctx *gin.Context) {
	shiftId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid shift ID"})
		return
	}

	shift, err := c.service.FindById(uint(shiftId))
	if err != nil {
		if err.Error() == "shift not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": shift})
}
