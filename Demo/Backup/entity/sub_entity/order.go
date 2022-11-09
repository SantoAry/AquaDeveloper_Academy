package sub_entity

//Object pada entity ini berhubungan dengan customer orders

//Static object Cart
type Cart struct {
	Cart_ID     int64         `json:"cart_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	User_Ref    int64         `json:"user_ref"` //Foreign key
	CartDetails []CartDetails `gorm:"constraint;foreignKey:Cart_Ref;references:Cart_ID"`
}

type CartDetails struct {
	CartDetails int64 `json:"cardetails_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Cart_Ref    int64 `json:"cart_ref"` //foreign key
	Order_Ref   int64 `json:"user_ref"` //foreign key

}

type Orders struct {
	Order_ID       int64         `json:"order_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Product_Ref    int64         `json:"product_ref"` //foreign key
	Quantity       int8          `json:"quantity"`
	Payment_Status bool          `json:"payment_status"`
	CartDetails    []CartDetails `gorm:"constraint;foreignKey:Order_Ref;references:Order_ID"`
	Invoice        []Invoice     `gorm:"constraint;foreignKey:InvoiceOrder_Ref;references:Order_ID"`
}

type Invoice struct {
	Invoice_ID       int64 `json:"invoice_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	InvoiceOrder_Ref int64 `json:"order_id"` //Foreign key
}
