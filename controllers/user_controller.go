package controllers

import (
	"attendance-app-api/dto"
	"attendance-app-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) IUserController {
	return &UserController{service: service}
}

func (c *UserController) FindAll(ctx *gin.Context) {
	users, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (c *UserController) FindById(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.service.FindById(uint(userId))
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserController) Create(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := c.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newUser})
}
