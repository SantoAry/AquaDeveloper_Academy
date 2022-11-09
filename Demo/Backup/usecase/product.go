package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"
	"strconv"

	"strings"

	"github.com/jinzhu/copier"
)

type IProductUseCase interface {
	CreateProduct(product response.CreateProductResponse) error
	GetProductList() ([]response.GetProductResponse, error)
	GetOneProduct(product_id int) (response.GetProductResponse, int64, error)
	SearchProducts(product_name string) ([]response.GetProductResponse, error)
	SortProducts(sort string) ([]response.GetProductResponse, error)
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

func (p ProductUsecase) GetOneProduct(id string) ([]response.GetProductResponse, int64, error) {
	product_id, _ := strconv.ParseInt(id, 0, 64)
	products, err := p.ProductRepository.GetOneProduct(product_id)
	if err != nil {
		return nil, 0, nil
	}
	OneProductResponse := []response.GetProductResponse{}
	copier.Copy(&OneProductResponse, &products)
	return OneProductResponse, product_id, nil
}

func (p ProductUsecase) SearchProducts(product_name string) ([]response.GetProductResponse, error) {
	products, err := p.ProductRepository.SearchProducts(product_name)
	if err != nil {
		return nil, err
	}

	filtered_products := []response.GetProductResponse{}
	copier.Copy(&filtered_products, &products) //Override dari variabel kanan ke kiri
	return filtered_products, nil
}

func (p ProductUsecase) SortProducts(sort string) ([]response.GetProductResponse, error) {
	sort = strings.ToLower(sort)
	products, err := p.ProductRepository.SortProducts(sort)
	if err != nil {
		return nil, err
	}

	sorted_products := []response.GetProductResponse{}
	copier.Copy(&sorted_products, &products) //Override dari variabel kanan ke kiri

	return sorted_products, nil
}
