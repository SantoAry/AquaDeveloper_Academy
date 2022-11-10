package entity

import "e-commerce/entity/sub_entity"

type Product struct {
	Product_ID   int64               `json:"product_id" gorm:"column:product_id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string              `json:"product_name"`
	Price        float64             `json:"price"`
	Orders       []sub_entity.Orders `gorm:"foreignKey:Product_Ref;references:Product_ID"`
}
