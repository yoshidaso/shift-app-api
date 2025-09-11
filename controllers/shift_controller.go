package controllers

import (
	"attendance-app-api/dto"
	"attendance-app-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IShiftController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

func (c *ShiftController) Create(ctx *gin.Context) {
	userName := ctx.Param("userName")
	var req dto.CreateShiftRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newShift, err := c.service.Create(req, userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newShift})
}

func (c *ShiftController) Delete(ctx *gin.Context) {
	shiftId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shift ID"})
		return
	}

	err = c.service.Delete(uint(shiftId))
	if err != nil {
		if err.Error() == "shift not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Shift deleted successfully"})
}
