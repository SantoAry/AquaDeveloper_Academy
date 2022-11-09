package entity

//Object pada entity ini berhubungan dengan customer orders

//Static object Cart
type Cart struct {
	Cart_ID int64 `json:"cart_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	User_ID int64 `json:"user_id"` //Foreign key
}

type CartDetails struct {
	CartDetails int64 `json:"cardetails_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Cart_ID     int64 `json:"cart_id"` //foreign key
	Order_ID    int64 `json:"user_id"` //foreign key
}

type Orders struct {
	Order_ID       int64 `json:"order_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Product_ID     int64 `json:"product_id"` //foreign key
	Quantity       int8  `json:"quantity"`
	Payment_Status bool  `json:"payment_status"`
}

type Invoice struct {
	Invoice_ID int64 `json:"invoice_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	// Apakah perlu? User_ID    int64 `json:"user_id"`
	Order_ID int64 `json:"order_id"` //Foreign key
}
