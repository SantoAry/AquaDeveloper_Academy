package config

import (
	"e-commerce/entity"
	"e-commerce/entity/sub_entity"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.Product{},
		&sub_entity.Cart{},
		&sub_entity.Orders{},
		&sub_entity.CartDetails{},
		&sub_entity.Invoice{},
	)
}
