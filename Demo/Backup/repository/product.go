package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	GetOneProduct(product_id int64) ([]entity.Product, error)
	SearchProducts(product_name string) ([]entity.Product, error)
	SortProducts(sort string) ([]entity.Product, error)
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

func (p ProductRepository) GetOneProduct(product_id int64) ([]entity.Product, error) {
	var product []entity.Product
	if err := p.db.Find(&product, product_id).Error; err != nil {
		return nil, nil
	}
	return product, nil
}

// By name
func (p ProductRepository) SearchProducts(product_name string) ([]entity.Product, error) {
	var filtered_products []entity.Product
	if err := p.db.Where("product_name LIKE ?", "%"+product_name+"%").Find(&filtered_products).Error; err != nil {
		return nil, err
	}
	return filtered_products, nil
}

// By Price Order
func (p ProductRepository) SortProducts(sort string) ([]entity.Product, error) {
	var sort_products []entity.Product
	var err error
	if sort == "ascending" || sort == "asc" {
		if err = p.db.Order("price asc").Find(&sort_products).Error; err != nil {
			return nil, err
		}
		return sort_products, nil
	}

	if sort == "descending" || sort == "desc" {
		if err = p.db.Order("price desc").Find(&sort_products).Error; err != nil {
			return nil, err
		}
		return sort_products, nil
	} else {
		if err = p.db.Order("price " + sort).Find(&sort_products).Error; err != nil {
			return nil, err
		}
		return sort_products, nil
	}
}
