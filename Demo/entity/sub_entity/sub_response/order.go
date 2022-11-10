package response

type OrderBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CartBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type InvoiceBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	User_ID int64       `json:"user_id"`
	Data    interface{} `json:"data"`
}

type GetCartResponse struct {
	Cart_ID  int64 `json:"cart_id" gorm:"column:cart_id;type:bigint;primaryKey;autoIncrement"`
	User_Ref int64 `json:"user_ref"` //Foreign key
}

type CartDetailsResponse struct {
	CartDetails int64 `json:"cartdetails_id" gorm:"column:cartdetails_id;type:bigint;primaryKey;autoIncrement"`
	Cart_Ref    int64 `json:"cart_ref"`  //foreign key
	Order_Ref   int64 `json:"order_ref"` //foreign key
}

type CreateOrderResponse struct {
	Order_ID       int64 `json:"order_id" gorm:"column:orders_id;type:bigint;primaryKey;autoIncrement"`
	Product_Ref    int64 `json:"product_ref"` //foreign key
	Quantity       int64 `json:"quantity"`
	Payment_Status bool  `json:"payment_status"`
}

type GetOrderResponse struct {
	Order_ID       int64 `json:"order_id" gorm:"column:orders_id;type:bigint;primaryKey;autoIncrement"`
	Product_Ref    int64 `json:"product_ref"` //foreign key
	Quantity       int64 `json:"quantity"`
	Payment_Status bool  `json:"payment_status"`
}

type GetInvoiceResponse struct {
	Invoice_ID       int64  `json:"invoice_id" gorm:"column:invoice_id;type:bigint;primaryKey;autoIncrement"`
	User_ID          int64  `json:"user_id"`
	User_name        string `json:"user_name"`
	InvoiceOrder_Ref int64  `json:"invoice_order_id"`
}
