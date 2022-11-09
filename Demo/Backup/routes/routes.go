// Define routes & endpoints
package routes

import (
	"e-commerce/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(echoApp *echo.Echo, userHandler *handler.UserHandler) {
	echoGroup := echoApp.Group("api/e-commerce/v1")

	echoGroup.GET("/users", userHandler.GetList)

	//To get data by ID
	echoGroup.GET("/users/:id", userHandler.GetOne)
	echoGroup.POST("/users", userHandler.CreateUser)

	//To update data by ID
	echoGroup.PUT("/users/:id", userHandler.UpdateData)

	//To delete data by ID
	echoGroup.DELETE("/users/:id", userHandler.DeleteData)

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
