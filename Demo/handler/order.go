package handler

import (
	response "e-commerce/entity/sub_entity/sub_response"
	"e-commerce/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

func NewOrderHandler(ordercase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: ordercase}
}

// Get list of all Users
func (o OrderHandler) GetAllCarts(ctx echo.Context) error {
	carts, err := o.orderUsecase.GetAllCarts()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.CartBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list cart failed",
			Data:    nil,
		})
	}

	if len(carts) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.CartBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no cart found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.CartBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got all cart",
		Data:    carts,
	})
}

// Create Order
func (o OrderHandler) CreateOrder(ctx echo.Context) error {
	user_id := ctx.QueryParam("user_id")
	createOrder := response.CreateOrderResponse{}
	//errcreated := o.orderUsecase.CreateOrder(createOrder, user_id)

	if errbind := ctx.Bind(&createOrder); errbind != nil {
		return ctx.JSON(http.StatusBadRequest, response.OrderBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if errcreated := o.orderUsecase.CreateOrder(createOrder, user_id); errcreated != nil {
		return ctx.JSON(http.StatusBadRequest, response.OrderBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to create order",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusCreated, response.OrderBaseResponse{
		Code:    http.StatusCreated,
		Message: "order added to cart successfully",
		Data:    nil,
	})
}

func (o OrderHandler) GetAllCartDetails(ctx echo.Context) error {
	cartdetails, err := o.orderUsecase.GetAllCartDetails()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.CartBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list cart details failed",
			Data:    nil,
		})
	}

	if len(cartdetails) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.CartBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no recorded order for carts",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.CartBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got all cart details",
		Data:    cartdetails,
	})
}

// Get all orders
func (o OrderHandler) GetAllOrders(ctx echo.Context) error {
	allorders, err := o.orderUsecase.GetAllOrders()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.CartBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list orders failed",
			Data:    nil,
		})
	}

	if len(allorders) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.CartBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no order has been made",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.CartBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got all orders",
		Data:    allorders,
	})
}

// Delete user by ID
func (o OrderHandler) DeleteOrder(ctx echo.Context) error {
	id := ctx.QueryParam("order_id")
	orders, err := o.orderUsecase.DeleteOrder(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.OrderBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to do payment",
			Data:    nil,
		})
	}

	if len(orders) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.OrderBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no order found, payment declined",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.OrderBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully deleted order with the following data",
		Data:    orders,
	})
}

func (o OrderHandler) GetInvoiceByID(ctx echo.Context) error {
	id := ctx.QueryParam("user_id")

	invoice, user_id, err := o.orderUsecase.GetInvoiceByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.InvoiceBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to get specified product data",
			User_ID: user_id,
			Data:    nil,
		})
	}

	if len(invoice) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.InvoiceBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no data found for specified product_id",
			User_ID: user_id,
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.InvoiceBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got specified product data",
		User_ID: user_id,
		Data:    invoice,
	})
}
