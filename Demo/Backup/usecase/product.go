package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type IProductUseCase interface {
	CreateProduct(product response.CreateProductResponse) error
	GetProductList() ([]response.GetProductResponse, error)
	GetOneProduct(product_id string) (response.GetProductResponse, error)
	FilterProducts(product_name string) ([]response.GetProductResponse, error)
}

type ProductUsecase struct {
	ProductRepository repository.IProductRepository
}

func NewProductUsecase(productRepository repository.IProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepository}
}

// Product
func (p ProductUsecase) CreateProduct(req response.CreateProductResponse) error {
	products := entity.Product{}
	copier.Copy(&products, &req)

	err := p.ProductRepository.CreateProduct(products)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductUsecase) GetProductList() ([]response.GetProductResponse, error) {
	products, err := p.ProductRepository.GetAllProducts()
	if err != nil {
		return nil, nil
	}

	productResponse := []response.GetProductResponse{}
	copier.Copy(&productResponse, &products) //Override dari variabel kanan ke kiri
	return productResponse, nil
}

func (p ProductUsecase) GetOneProduct(product_id string) ([]response.GetProductResponse, error) {
	products, err := p.ProductRepository.GetOneProduct(product_id)
	if err != nil {
		return nil, nil
	}
	OneProductResponse := []response.GetProductResponse{}
	copier.Copy(&OneProductResponse, &products)
	return OneProductResponse, nil
}

func (p ProductUsecase) FilterProducts(product_name string) ([]response.GetProductResponse, error) {
	filtered_products, err := p.ProductRepository.FilterProducts(product_name)
	if err != nil {
		return nil, nil
	}

	productResponse := []response.GetProductResponse{}
	copier.Copy(&productResponse, &filtered_products) //Override dari variabel kanan ke kiri
	return productResponse, nil
}
