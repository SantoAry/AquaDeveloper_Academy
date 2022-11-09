// Repository interacts with Database
package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	//Role
	CreateRole(role entity.Role) error
	GetAllRoles() ([]entity.Role, error)

	//User
	CreateUser(user entity.User) error
	GetAllUsers() ([]entity.User, error)
	GetOneUser(id string) ([]entity.User, error)
	UpdateUser(id string, user entity.User) error
	DeleteUser(id string) ([]entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Role
func (u UserRepository) CreateRole(role entity.Role) error {
	if err := u.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetAllRoles() ([]entity.Role, error) {
	var roles []entity.Role
	if err := u.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// Users
func (u UserRepository) CreateUser(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Definisikan GetOne dengan ctx.Params dari handler id untuk mencari data dengan id integer pada /users/:id
func (u UserRepository) GetOneUser(id string) ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Find(&users, id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Definisikan Update untuk melakukan cek bila update gagal dari value u.db.Updates(&user).Error
func (u UserRepository) UpdateUser(id string, user entity.User) error {
	if err := u.db.Model(user).Where("user_id = ?", id).Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// Definisikan DeleteUser dengan ctx.Params dari handler id untuk mencari bila data yang ingin di delete ada atau tidak
func (u UserRepository) DeleteUser(id string) ([]entity.User, error) {
	var users []entity.User
	//subquery := u.db.Delete(&users, id)

	if err := u.db.Clauses(clause.Returning{}).Delete(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}
