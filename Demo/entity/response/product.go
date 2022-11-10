package response

type ProductBaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProductIDResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Product_ID int64       `json:"product_id"`
	Data       interface{} `json:"data"`
}

type ProductSearchResponse struct {
	Code          int         `json:"code"`
	Message       string      `json:"message"`
	Search_Filter string      `json:"search_filter"`
	Data          interface{} `json:"data"`
}

type ProductSortResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Sort_Type string      `json:"sort_type"`
	Data      interface{} `json:"data"`
}

type CreateProductResponse struct {
	Product_ID   int64   `json:"product_id" gorm:"column:product_id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string  `json:"product_name"`
	Price        float64 `json:"price"`
	//Orders       []sub_entity.Orders `gorm:"foreignKey:Product_Ref;references:Product_ID"`
}

// Redundant, make into get response for one?
type GetProductResponse struct {
	Product_ID   int64   `json:"product_id" gorm:"column:product_id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string  `json:"product_name"`
	Price        float64 `json:"price"`
	//Orders       []sub_entity.Orders `gorm:"foreignKey:Product_Ref;references:Product_ID"`
}

type FilterProductResponse struct {
	Product_ID   int64   `json:"product_id" gorm:"column:product_id;type:bigint;primaryKey;autoIncrement"`
	Product_Name string  `json:"product_name"`
	Price        float64 `json:"price"`
	//Orders       []sub_entity.Orders `gorm:"foreignKey:Product_Ref;references:Product_ID"`
}
