// Usecase for handler request
package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type IUserUseCase interface {
	//Role
	CreateRole(role response.CreateRoleResponse) error
	GetAllRoles() ([]response.GetRoleResponse, error)

	//User
	CreateUser(user response.CreateUserResponse) error
	GetAllUsers() ([]response.GetUserResponse, error)
	GetOneUser(id string) (response.GetUserResponse, error)
	UpdateUser(id string, user entity.User) error
	DeleteUser(id string) (response.DeleteUserResponse, error)
}

type UserUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

// Roles
func (u UserUsecase) CreateRole(req response.CreateRoleResponse) error {
	role := entity.Role{}
	copier.Copy(&role, &req)

	err := u.UserRepository.CreateRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetAllRoles() ([]response.GetRoleResponse, error) {
	roles, err := u.UserRepository.GetAllRoles()
	if err != nil {
		return nil, err
	}

	roleResponse := []response.GetRoleResponse{}
	copier.Copy(&roleResponse, &roles) //Override dari variabel kanan ke kiri
	return roleResponse, nil

}

// Users
func (u UserUsecase) CreateUser(req response.CreateUserResponse) error {
	users := entity.User{}
	copier.Copy(&users, &req)

	err := u.UserRepository.CreateUser(users)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetAllUsers() ([]response.GetUserResponse, error) {
	users, err := u.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users) //Override dari variabel kanan ke kiri
	return userResponse, nil

}

// Usecase for GetOne
func (u UserUsecase) GetOneUser(id string) ([]response.GetUserResponse, error) {
	users, err := u.UserRepository.GetOneUser(id)
	if err != nil {
		return nil, err
	}
	OneUserResponse := []response.GetUserResponse{}
	copier.Copy(&OneUserResponse, &users)
	return OneUserResponse, nil
}

// Usecase for Update
func (u UserUsecase) UpdateUser(id string, req response.UpdateUserResponse) error {
	users := entity.User{}
	copier.Copy(&users, &req)

	err := u.UserRepository.UpdateUser(id, users)
	if err != nil {
		return err
	}

	return nil
}

// Usecase for DeleteUser
func (u UserUsecase) DeleteUser(id string) ([]entity.User, error) {
	users, err := u.UserRepository.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return users, nil
}
