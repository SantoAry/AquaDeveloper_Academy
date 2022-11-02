package main

import (
	"fiber/config"
	"fiber/handler"
	"fiber/repository"
	"fiber/routes"
	"fiber/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := fiber.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	routes.Routes(app, userHandler)

	//Listening to localhost:8080
	app.Listen(":8080")
}
