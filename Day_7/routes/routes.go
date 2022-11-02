package routes

import (
	"fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("api/v1")
	r.Get("/users", userHandler.GetList)

	//To get data by ID
	r.Get("/users/:id", userHandler.GetOne)
	r.Post("/users", userHandler.CreateUser)

	//To update data by ID
	r.Put("/users/:id", userHandler.UpdateData)

	//To delete data by ID
	r.Delete("/users/:id", userHandler.DeleteData)
}
