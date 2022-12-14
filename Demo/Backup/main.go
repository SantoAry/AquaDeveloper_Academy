// API entry point
package main

import (
	"e-commerce/config"
	"e-commerce/handler"
	"e-commerce/repository"
	"e-commerce/routes"
	"e-commerce/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.AutoMigrate()

	echoApp := echo.New()

	//User & Role
	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	routes.UserRoutes(echoApp, userHandler)

	//Product
	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	routes.ProductRoutes(echoApp, productHandler)

	//Order; Cart; Cart Details; Invoice
	orderRepository := repository.NewOrderRepository(config.DB)
	orderUsecase := usecase.NewOrderUsecase(orderRepository)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	routes.OrderRoutes(echoApp, orderHandler)

	//Listening to localhost:8080
	echoApp.Logger.Fatal(echoApp.Start(":9000"))
}
