package main

import (
	"attendance-app-api/controllers"
	"attendance-app-api/infra"
	"attendance-app-api/repositories"
	"attendance-app-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetUpDB()

	shiftRepository := repositories.NewShiftRepository(db)
	shiftService := services.NewShiftService(shiftRepository)
	shiftController := controllers.NewShiftController(shiftService)

	r := gin.Default()
	r.GET("/shifts", shiftController.FindAll)
	r.GET("/shifts/:id", shiftController.FindById)
	r.POST("/shifts", shiftController.Create)
	r.Run(":8080")
}
