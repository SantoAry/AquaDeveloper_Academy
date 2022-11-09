// Define routes & endpoints
package routes

import (
	"e-commerce/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(echoApp *echo.Echo, userHandler *handler.UserHandler) {
	echoGroup := echoApp.Group("api/e-commerce/v1")

	//Insert user data
	echoGroup.POST("/users", userHandler.CreateUser)

	//Get all user data
	echoGroup.GET("/users", userHandler.GetAllUsers)

	//To get user data by ID
	echoGroup.GET("/users/:id", userHandler.GetOneUser)

	//To check if empty, then update data by ID (Only update non Primary Key allowed)
	echoGroup.PUT("/users/:id", userHandler.UpdateUser)

	//For existing data, update data by ID
	echoGroup.PATCH("/users/:id", userHandler.UpdateUser)

	//To delete data by ID
	echoGroup.DELETE("/users/:id", userHandler.DeleteUser)

	//Create role
	echoGroup.POST("/roles", userHandler.CreateRole)

	//Get all roles
	echoGroup.GET("/roles", userHandler.GetAllRoles)

	//Insert Cart
	//r.Post("/cart_id", ....)

	//List users
	//r.Get("/users", ....)

	//List users by id
	//r.Get("/users/:user_id", ...)

	//Insert Users
	//r.Post("/users", userHandler.CreateUser)

	//Insert cart for user
	//r.Post("/users/:user_id/:cart_id", )

	//To update data by ID
	//r.Put("/users/:id", userHandler.UpdateData)

	//To delete data by ID
	//r.Delete("/users/:id", userHandler.DeleteData)
}

func ProductRoutes(echoApp *echo.Echo, productHandler *handler.ProductHandler) {
	echoGroup := echoApp.Group("api/e-commerce/v1")
	echoGroup.GET("/products", productHandler.GetProductList)

	//To get data by ID
	echoGroup.GET("/products/:product_id", productHandler.GetOneProduct)

	//Filter
	//Search by name product
	echoGroup.GET("/products/search", productHandler.SearchProducts)

	//Order by price (descending and ascending)
	echoGroup.GET("/products/sort", productHandler.SortProducts)

	//Insert Product
	echoGroup.POST("/products", productHandler.CreateProduct)
}
