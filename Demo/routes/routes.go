// Define routes & endpoints
package routes

import (
	"crypto/subtle"
	"e-commerce/handler"
	"e-commerce/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(echoApp *echo.Echo, userHandler *handler.UserHandler) {
	echoUser := echoApp.Group("api/e-commerce/v1")
	echoAdmin := echoApp.Group("api/e-commerce/v1/admin")

	//Basic authentication for admin
	echoAdmin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("efishery")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	//Create role
	echoAdmin.POST("/roles", userHandler.CreateRole)

	//Get all roles
	echoAdmin.GET("/roles", userHandler.GetAllRoles)

	//Create user data -> Automatically create a record in Cart table
	echoUser.POST("/users", userHandler.CreateUser)

	//Get all user data or by roles with query param ->	/users?role=Admin; Customer; Merchant
	echoUser.GET("/users", userHandler.GetAllUsers)

	//To get user data by ID
	echoUser.GET("/users/:id", userHandler.GetOneUser)

	//Update user data by ID (foreign key Role_Ref must be specified) Update only non-constrained parts
	echoUser.PUT("/users/:id", userHandler.UpdateUser)

	//To delete data by ID
	echoAdmin.DELETE("/users/:id", userHandler.DeleteUser)

}

func ProductRoutes(echoApp *echo.Echo, productHandler *handler.ProductHandler) {
	echoGroup := echoApp.Group("api/e-commerce/v1")
	echoMerchant := echoApp.Group("api/e-commerce/v1/merchant")

	//Basic authentication for merchant to post item
	echoMerchant.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("merchant")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("efishery")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	//Insert Product
	echoMerchant.POST("/products", productHandler.CreateProduct)

	//Get all Products
	echoGroup.GET("/products", productHandler.GetProductList)

	//Filter
	//To get product by ID
	echoGroup.GET("/products/:product_id", productHandler.GetOneProduct)

	//Search by name product can be sorted as well
	echoGroup.GET("/products/search", productHandler.SearchProducts)

	//Order by price (descending and ascending)
	echoGroup.GET("/products/sort", productHandler.SortProducts)

	//Order by price (descending and ascending)
	echoGroup.GET("/products/filter", productHandler.SortProductsPrice)
}

func OrderRoutes(echoApp *echo.Echo, orderHandler *handler.OrderHandler) {
	echoGroup := echoApp.Group("api/e-commerce/v1")

	//List all Carts
	echoGroup.GET("/carts", orderHandler.GetAllCarts)

	//User make order
	echoGroup.POST("/users/orders", orderHandler.CreateOrder)

	//Get all orders
	echoGroup.GET("/orders", orderHandler.GetAllOrders)

	//List all Cart Details
	echoGroup.GET("/carts/cartdetails", orderHandler.GetAllCartDetails)

	//Pay for orders -> Remove item from cart
	echoGroup.DELETE("/users/payment", orderHandler.DeleteOrder)

	//Get invoice by ID
	echoGroup.GET("/invoice", orderHandler.GetInvoiceByID)
}
