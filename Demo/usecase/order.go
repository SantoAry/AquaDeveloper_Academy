package usecase

import (
	"e-commerce/entity/sub_entity"
	response "e-commerce/entity/sub_entity/sub_response"
	"e-commerce/repository"
	"strconv"

	"github.com/jinzhu/copier"
)

type IOrderUseCase interface {
	//Invoice
	GetInvoiceByID(id string) ([]response.GetInvoiceResponse, int64, error)

	//Cart
	//Cart creation is made from user creation
	GetAllCarts() ([]response.GetCartResponse, error)

	//Order
	CreateOrder(req response.GetOrderResponse, user_id string) error
	GetAllOrders() ([]response.GetOrderResponse, error)
	DeleteOrder(id string) ([]sub_entity.Orders, error)

	//Cart Details
	GetAllCartDetails(user_id string) ([]response.CartDetailsResponse, error)
}

type OrderUsecase struct {
	OrderRepository repository.IOrderRepository
}

func NewOrderUsecase(orderRepository repository.IOrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepository}
}

// List all carts
func (o OrderUsecase) GetAllCarts() ([]response.GetCartResponse, error) {
	carts, err := o.OrderRepository.GetAllCarts()
	if err != nil {
		return nil, err
	}

	cartResponse := []response.GetCartResponse{}
	copier.Copy(&cartResponse, &carts) //Override dari variabel kanan ke kiri
	return cartResponse, nil
}

// Create Order
func (o OrderUsecase) CreateOrder(req response.CreateOrderResponse, user_id string) error {
	order := sub_entity.Orders{}
	copier.Copy(&order, &req)

	err := o.OrderRepository.CreateOrder(order, user_id)
	if err != nil {
		return err
	}
	return nil
}

// Get all orders
func (o OrderUsecase) GetAllOrders() ([]response.GetOrderResponse, error) {
	orders, err := o.OrderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}

	allorders := []response.GetOrderResponse{}
	copier.Copy(&allorders, &orders) //Override dari variabel kanan ke kiri
	return allorders, nil
}

func (o OrderUsecase) GetAllCartDetails(user_id string) ([]response.CartDetailsResponse, error) {
	id, _ := strconv.ParseInt(user_id, 0, 64)
	cartdetails, err := o.OrderRepository.GetAllCartDetails(id)
	if err != nil {
		return nil, err
	}

	cartResponse := []response.CartDetailsResponse{}
	copier.Copy(&cartResponse, &cartdetails) //Override dari variabel kanan ke kiri
	return cartResponse, nil
}

// Usecase for DeleteOrder
func (o OrderUsecase) DeleteOrder(id string) ([]sub_entity.Orders, error) {
	orders, err := o.OrderRepository.DeleteOrder(id)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// Get invoice by ID
func (o OrderUsecase) GetInvoiceByID(id string) ([]response.GetInvoiceResponse, int64, error) {
	user_id, _ := strconv.ParseInt(id, 0, 64)
	invoice, err := o.OrderRepository.GetInvoiceByID(id)
	if err != nil {
		return nil, 0, nil
	}
	InvoiceID := []response.GetInvoiceResponse{}
	copier.Copy(&InvoiceID, &invoice)
	return InvoiceID, user_id, nil
}
