package sub_entity

//Object pada entity ini berhubungan dengan customer orders

//Static object Cart
type Cart struct {
	Cart_ID     int64         `json:"cart_id" gorm:"column:cart_id;type:bigint;primaryKey;autoIncrement"`
	User_Ref    int64         `json:"user_ref"` //Foreign key
	CartDetails []CartDetails `gorm:"foreignKey:Cart_Ref;references:Cart_ID;constraint:OnUpdate:SETNULL,OnDelete:SET NULL"`
}

type CartDetails struct {
	CartDetails int64 `json:"cartdetails_id" gorm:"column:cartdetails_id;type:bigint;primaryKey;autoIncrement"`
	Cart_Ref    int64 `json:"cart_ref"`  //foreign key
	Order_Ref   int64 `json:"order_ref"` //foreign key
}

type Orders struct {
	Order_ID       int64         `json:"order_id" gorm:"column:orders_id;type:bigint;primaryKey;autoIncrement"`
	Product_Ref    int64         `json:"product_ref"` //foreign key
	Quantity       int64         `json:"quantity"`
	Payment_Status bool          `json:"payment_status"`
	CartDetails    []CartDetails `gorm:"foreignKey:Order_Ref;references:Order_ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Invoice struct {
	Invoice_ID       int64  `json:"invoice_id" gorm:"column:invoice_id;type:bigint;primaryKey;autoIncrement"`
	User_ID          int64  `json:"user_id"`
	User_Name        string `json:"user_name"`
	InvoiceOrder_Ref int64  `json:"invoice_order_id"`
}
