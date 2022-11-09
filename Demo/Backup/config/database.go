package config

import (
	"os"

	"e-commerce/entity"
	"e-commerce/entity/sub_entity"

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
		&entity.User{},
		&entity.Role{},
		&entity.Product{},
		&sub_entity.Cart{},
		&sub_entity.CartDetails{},
		&sub_entity.Orders{},
		&sub_entity.Invoice{},
	)
}
