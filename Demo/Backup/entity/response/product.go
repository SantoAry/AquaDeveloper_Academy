package response

type ProductBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateProductResponse struct {
	Product_ID   int64   `json:"product_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string  `json:"product_name"`
	Price        float64 `json:"price"`
	//Quantity     int     `json:"quantity"`
}

// Redundant, make into get response for one?
type GetProductResponse struct {
	Product_ID   int64   `json:"product_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string  `json:"product_name"`
	Price        float64 `json:"price"`
	//Quantity     int     `json:"quantity"`
}

type FilterProductResponse struct {
	Product_ID   int64   `json:"product_id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string  `json:"product_name"`
	Price        float64 `json:"price"`
	//Quantity     int     `json:"quantity"`
}
