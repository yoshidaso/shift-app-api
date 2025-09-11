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

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	r := gin.Default()
	r.GET("/shifts", shiftController.FindAll)
	r.GET("/shifts/:id", shiftController.FindById)
	r.POST("/shifts/:userName", shiftController.Create)
	r.DELETE("/shifts/:id", shiftController.Delete)

	r.GET("/users", userController.FindAll)
	r.GET("/users/:id", userController.FindById)
	r.POST("/users", userController.Create)
	r.DELETE("/users/:id", userController.Delete)
	r.Run(":8080")
}
