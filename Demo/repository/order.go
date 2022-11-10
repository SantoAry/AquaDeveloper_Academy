package repository

import (
	"e-commerce/entity"
	"e-commerce/entity/sub_entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IOrderRepository interface {
	//Invoice
	GetInvoiceByID(user_id string) ([]sub_entity.Invoice, error)

	//Cart
	//Cart creation -> From User creation
	GetAllCarts() ([]sub_entity.Cart, error)

	//Order
	CreateOrder(order sub_entity.Orders, user_id string) error
	GetAllOrders() ([]sub_entity.Orders, error)
	DeleteOrder(id string) ([]sub_entity.Orders, error)

	//Cart Details
	GetAllCartDetails() ([]sub_entity.CartDetails, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Get all cart
func (o OrderRepository) GetAllCarts() ([]sub_entity.Cart, error) {
	var carts []sub_entity.Cart
	if err := o.db.Find(&carts).Error; err != nil {
		return nil, nil
	}
	return carts, nil
}

// Create order
func (o OrderRepository) CreateOrder(order sub_entity.Orders, user_id string) error {
	var carts sub_entity.Cart
	var new_cart_details sub_entity.CartDetails

	if err := o.db.Create(&order).Error; err != nil {
		return err
	}

	if err := o.db.Joins("INNER JOIN users on users.user_id=carts.user_ref").Group("carts.cart_id").Where("users.user_id=?", user_id).Find(&carts).Error; err != nil {
		return err
	}

	if err := o.db.Joins("INNER JOIN orders on orders.order_id=carts_details.orders_ref").Group("cart_details.cartdetails_id").Error; err != nil {
		return err
	}

	new_cart_details = sub_entity.CartDetails{Cart_Ref: carts.Cart_ID, Order_Ref: order.Order_ID}
	if err := o.db.Create(&new_cart_details).Error; err != nil {
		return err
	}

	return nil
}

// Get all orders
func (o OrderRepository) GetAllOrders() ([]sub_entity.Orders, error) {
	var allorders []sub_entity.Orders
	if err := o.db.Find(&allorders).Error; err != nil {
		return nil, err
	}
	return allorders, nil
}

// Get all cart details
func (o OrderRepository) GetAllCartDetails() ([]sub_entity.CartDetails, error) {
	var cartdetails []sub_entity.CartDetails
	if err := o.db.Find(&cartdetails).Error; err != nil {
		return nil, err
	}
	return cartdetails, nil
}

// Delete order
func (o OrderRepository) DeleteOrder(id string) ([]sub_entity.Orders, error) {
	var orders []sub_entity.Orders
	var invoice sub_entity.Invoice
	var cartdetails sub_entity.CartDetails
	var cart sub_entity.Cart
	var user entity.User

	if err := o.db.Model(&orders).Where("orders_id = ?", id).Update("payment_status", true).Error; err != nil {
		return nil, err
	}

	//Join orders with cart details
	if err := o.db.Joins("INNER JOIN orders on orders.orders_id=cart_details.order_ref").Group("cart_details.cartdetails_id").Where("orders.orders_id=?", id).Find(&cartdetails).Error; err != nil {
		return nil, err
	}

	//Join cart details with cart
	if err := o.db.Joins("INNER JOIN cart_details on cart_details.cart_ref=carts.cart_id").Group("carts.cart_id").Where("cart_details.cart_ref=?", cartdetails.Cart_Ref).Find(&cart).Error; err != nil {
		return nil, err
	}

	//Join cart with user
	if err := o.db.Joins("INNER JOIN carts on carts.user_ref=users.user_id").Group("users.user_id").Where("carts.user_ref=?", cart.User_Ref).Find(&user).Error; err != nil {
		return nil, err
	}

	//Create invoice
	invoice = sub_entity.Invoice{User_ID: user.User_ID, User_Name: user.Name, InvoiceOrder_Ref: cartdetails.Order_Ref}
	if err := o.db.Create(&invoice).Error; err != nil {
		return nil, err
	}

	//Delete order record and remove cart
	if err := o.db.Clauses(clause.Returning{}).Delete(&orders, id).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// Get invoice by ID
func (o OrderRepository) GetInvoiceByID(id string) ([]sub_entity.Invoice, error) {
	var invoice []sub_entity.Invoice
	var user entity.User

	if err := o.db.Find(&user, id).Error; err != nil {
		return nil, err
	}

	if err := o.db.Where("user_id = ?", user.User_ID).Find(&invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}
