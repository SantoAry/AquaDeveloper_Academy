package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	GetOneProduct(product_id string) ([]entity.Product, error)
	FilterProducts(product_name string) ([]entity.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Buat product
func (p ProductRepository) CreateProduct(product entity.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

// Get all product
func (p ProductRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, nil
	}
	return products, nil
}

func (p ProductRepository) GetOneProduct(product_id string) ([]entity.Product, error) {
	var products []entity.Product
	if err := p.db.Find(&products, product_id).Error; err != nil {
		return nil, nil
	}
	return products, nil
}

// By name
func (p ProductRepository) FilterProducts(product_name string) ([]entity.Product, error) {
	var filtered_products []entity.Product
	if err := p.db.Where("product_name LIKE ?", "%"+product_name).Find(&filtered_products).Error; err != nil {
		return nil, nil
	}
	return filtered_products, nil
}

//p.db.Where("product_name <> ?", product_name).Find(&filtered_products)
