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
	product_id := ctx.Param("product_id")
	products, err := h.productUsecase.GetOneProduct(product_id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to get specified product data",
			Data:    nil,
		})
	}

	if len(products) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no data found for specified product_id",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got specified product data",
		Data:    products,
	})
}

// By name
func (h ProductHandler) FilterProduct(ctx echo.Context) error {
	product_name := ctx.QueryParam("name")
	filtered_products, err := h.productUsecase.FilterProducts(product_name)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to get specified product data",
			Data:    nil,
		})
	}

	if len(filtered_products) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no data found for specified filter",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got filter product data",
		Data:    filtered_products,
	})
}
