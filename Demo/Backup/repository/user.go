// Repository interacts with Database
package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	GetOne(id string) ([]entity.User, error)
	Update(id string, user entity.User) error
	DeleteUser(id string) ([]entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

// Definisikan GetOne dengan ctx.Params dari handler id untuk mencari data dengan id integer pada /users/:id
func (u UserRepository) GetOne(id string) ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Find(&users, id).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

// Definisikan Update untuk melakukan cek bila update gagal dari value u.db.Updates(&user).Error
func (u UserRepository) Update(id string, user entity.User) error {
	if err := u.db.Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// Definisikan DeleteUser dengan ctx.Params dari handler id untuk mencari bila data yang ingin di delete ada atau tidak
func (u UserRepository) DeleteUser(id string) ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Find(&users, id).Error; err != nil {
		return nil, nil
	}
	return users, nil
}
