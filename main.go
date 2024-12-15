package main

import (
	"attendance-app-api/controllers"
	"attendance-app-api/models"
	"attendance-app-api/repositories"
	"attendance-app-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	shifts := []models.Shift{
		{ID: 1, UserID: 1, StartTime: "08:00", EndTime: "17:00"},
		{ID: 2, UserID: 2, StartTime: "09:00", EndTime: "18:00"},
		{ID: 3, UserID: 3, StartTime: "10:00", EndTime: "19:00"},
	}

	shiftRepository := repositories.NewShiftMemoryRepository(shifts)
	shiftService := services.NewShiftService(shiftRepository)
	shiftController := controllers.NewShiftController(shiftService)

	r := gin.Default()
	r.GET("/shifts", shiftController.FindAll)
	r.GET("/shifts/:id", shiftController.FindById)
	r.Run("localhost:8080")
}
