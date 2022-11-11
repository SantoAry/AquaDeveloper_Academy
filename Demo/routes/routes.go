// Define routes & endpoints
package routes

import (
	"crypto/subtle"
	"e-commerce/handler"
	"e-commerce/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(echoApp *echo.Echo, userHandler *handler.UserHandler) {
	echoGroup := echoApp.Group("api/e-commerce/v1")
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
	echoGroup.POST("/users", userHandler.CreateUser)

	//Get all user data or by roles with query param ->	/users?role=Admin; Customer; Merchant
	//Need admin authorization, if need to open to all access change echoAdmin to echoUser
	echoAdmin.GET("/users", userHandler.GetAllUsers)

	//To get user data by ID
	echoGroup.GET("/users/:id", userHandler.GetOneUser)

	//Update user data by ID (foreign key Role_Ref must be specified) Update only non-constrained parts
	echoGroup.PUT("/users/:id", userHandler.UpdateUser)

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
	echoAdmin := echoApp.Group("api/e-commerce/v1/admin")
	echoUser := echoApp.Group("api/e-commerce/v1/users")

	//Basic authentication for admin
	echoAdmin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("efishery")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	//Basic authentication for users
	echoUser.Use(middleware.BasicAuth(func(user_id, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(user_id), []byte(user_id)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("efishery")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	//List all Carts
	echoAdmin.GET("/carts", orderHandler.GetAllCarts)

	//User make order
	echoUser.POST("/user/orders", orderHandler.CreateOrder)

	//Get all orders
	echoAdmin.GET("/active_order", orderHandler.GetAllOrders)

	//List all Cart Details
	echoUser.GET("/carts/cartdetails", orderHandler.GetAllCartDetails)

	//Pay for orders -> Remove item from cart
	echoUser.DELETE("/payment", orderHandler.DeleteOrder)

	//Get invoice by ID
	echoUser.GET("/invoice", orderHandler.GetInvoiceByID)
}
