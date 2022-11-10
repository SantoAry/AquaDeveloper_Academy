// Repository interacts with Database
package repository

import (
	"e-commerce/entity"
	"e-commerce/entity/sub_entity"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	//Role
	CreateRole(role entity.Role) error
	GetAllRoles() ([]entity.Role, error)

	//User
	CreateUser(user entity.User) error
	GetAllUsers(role string) ([]entity.User, error)
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

	if user.Role_Ref == 2 || user.Role_Ref == 3 {
		new_cart := sub_entity.Cart{User_Ref: user.User_ID}
		if err := u.db.Create(&new_cart).Error; err != nil {
			return err
		}
		if err := u.db.Model(&user).Update("Cart_ID", new_cart.Cart_ID).Error; err != nil {
			return err
		}
		return nil
	} else {
		if err := u.db.Model(&user).Update("Cart_ID", 0).Error; err != nil {
			return err
		}
		return nil
	}
}

func (u UserRepository) GetAllUsers(role string) ([]entity.User, error) {
	var users []entity.User
	if role == "Admin" || role == "Customer" || role == "Merchant" {
		if err := u.db.Joins("INNER JOIN roles on roles.role_id=users.role_ref").Group("users.user_id").Where("roles.role_type=?", role).Find(&users).Error; err != nil {
			return nil, err
		}
		return users, nil
	} else if role == "" {
		if err := u.db.Find(&users).Error; err != nil {
			return nil, err
		}
		return users, nil
	} else {
		if err := u.db.Joins("INNER JOIN roles on roles.role_id=users.role_ref").Group("users.user_id").Where("roles.role_type=?", role).Find(&users).Error; err != nil {
			return nil, err
		}
		return users, nil
	}
}

// Definisikan GetOne dengan ctx.Params dari handler id untuk mencari data dengan id integer pada /users/:id
func (u UserRepository) GetOneUser(id string) ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Find(&users, id).Error; err != nil {
		return nil, err
	}
	// if err = u.db.Joins("JOIN roles on roles.role_id=users.role_ref").Group("artists.id").Find(&users).Error;
	return users, nil
}

// Definisikan Update untuk melakukan cek bila update gagal dari value u.db.Updates(&user).Error
func (u UserRepository) UpdateUser(id string, user entity.User) error {

	var users entity.User
	if err := u.db.Find(&users, id).Error; err != nil {
		return err
	}

	if users.Role_Ref == 2 || users.Role_Ref == 3 {

		if err := u.db.Model(user).Where("user_id = ?", id).Updates(&user).Error; err != nil {
			return err
		}
	} else if users.Role_Ref == 1 {
		return fmt.Errorf("Unauthorized")
	}

	return nil

}

// Definisikan DeleteUser dengan ctx.Params dari handler id untuk mencari bila data yang ingin di delete ada atau tidak
func (u UserRepository) DeleteUser(id string) ([]entity.User, error) {
	var users []entity.User

	if err := u.db.Clauses(clause.Returning{}).Delete(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}
