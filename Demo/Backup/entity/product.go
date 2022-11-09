package entity

type Product struct {
	Product_ID   int64   `json:"product_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string  `json:"product_name"`
	Price        float64 `json:"price"`
	//Quantity     int     `json:"quantity"`
}
