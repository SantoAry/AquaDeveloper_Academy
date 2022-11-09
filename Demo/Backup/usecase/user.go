// Usecase for handler request
package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type IUserUseCase interface {
	CreateUser(user response.CreateUserResponse) error
	GetListUser() ([]response.GetUserResponse, error)
	GetOne(id string) (response.GetUserResponse, error)
	Update(id string, user entity.User) error
	DeleteUser(id string) (response.DeleteUserResponse, error)
}

type UserUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (u UserUsecase) CreateUser(req response.CreateUserResponse) error {
	users := entity.User{}
	copier.Copy(&users, &req)

	err := u.UserRepository.Create(users)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetList() ([]response.GetUserResponse, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		return nil, nil
	}

	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users) //Override dari variabel kanan ke kiri
	return userResponse, nil

}

// Usecase for GetOne
func (u UserUsecase) GetOne(id string) ([]response.GetUserResponse, error) {
	users, err := u.UserRepository.GetOne(id)
	if err != nil {
		return nil, nil
	}
	OneUserResponse := []response.GetUserResponse{}
	copier.Copy(&OneUserResponse, &users)
	return OneUserResponse, nil
}

// Usecase for Update
func (u UserUsecase) Update(id string, req response.UpdateUserResponse) error {
	users := entity.User{}
	copier.Copy(&users, &req)

	err := u.UserRepository.Update(id, users)
	if err != nil {
		return err
	}

	return nil
}

// Usecase for DeleteUser
func (u UserUsecase) DeleteUser(id string) ([]entity.User, error) {
	users, err := u.UserRepository.DeleteUser(id)
	if err != nil {
		return nil, nil
	}

	return users, nil
}
