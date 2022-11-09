package handler

import (
	"e-commerce/entity/response"
	"e-commerce/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productcase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productcase}
}

// Product
func (h ProductHandler) CreateProduct(ctx echo.Context) error {
	createProduct := response.CreateProductResponse{}
	if err := ctx.Bind(&createProduct); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ProductBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.productUsecase.CreateProduct(createProduct); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ProductBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to create product",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusCreated, response.ProductBaseResponse{
		Code:    http.StatusCreated,
		Message: "product created successfully",
		Data:    nil,
	})
}

func (h ProductHandler) GetProductList(ctx echo.Context) error {
	products, err := h.productUsecase.GetProductList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ProductBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list products failed",
			Data:    nil,
		})
	}

	if len(products) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.ProductBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no products found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.ProductBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got all products",
		Data:    products,
	})
}

func (h ProductHandler) GetOneProduct(ctx echo.Context) error {
	id := ctx.Param("product_id")

	products, product_id, err := h.productUsecase.GetOneProduct(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ProductIDResponse{
			Code:       http.StatusBadRequest,
			Message:    "failed to get specified product data",
			Product_ID: product_id,
			Data:       nil,
		})
	}

	if len(products) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.ProductIDResponse{
			Code:       http.StatusBadRequest,
			Message:    "no data found for specified product_id",
			Product_ID: product_id,
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.ProductIDResponse{
		Code:       http.StatusOK,
		Message:    "successfully got specified product data",
		Product_ID: product_id,
		Data:       products,
	})
}

// By name
func (h ProductHandler) SearchProducts(ctx echo.Context) error {
	product_name := ctx.QueryParam("name")
	filtered_products, err := h.productUsecase.SearchProducts(product_name)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ProductFilterResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to filter products",
			Data:    nil,
		})
	}

	if len(filtered_products) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.ProductFilterResponse{
			Code:    http.StatusBadRequest,
			Message: "no data found for this item",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.ProductFilterResponse{
		Code:    http.StatusOK,
		Message: "successfully filtered products",
		Data:    filtered_products,
	})
}

// order
func (h ProductHandler) SortProducts(ctx echo.Context) error {
	sort := ctx.QueryParam("sort_price")
	sorted_products, err := h.productUsecase.SortProducts(sort)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ProductSortResponse{
			Code:      http.StatusBadRequest,
			Message:   "sort type invalid (available sort type:ascending; asc; descending; desc)",
			Sort_Type: sort,
			Data:      nil,
		})
	}
	//need to test when theres nothing
	if len(sorted_products) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.ProductSortResponse{
			Code:      http.StatusBadRequest,
			Message:   "no products are available to sort",
			Sort_Type: sort,
			Data:      nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.ProductSortResponse{
		Code:      http.StatusOK,
		Message:   "successfully sorted product data",
		Sort_Type: sort,
		Data:      sorted_products,
	})
}
