// Handles response to client
package handler

import (
	"e-commerce/config"
	"e-commerce/entity/response"
	"e-commerce/usecase"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usercase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: usercase}
}

func (h UserHandler) CreateUser(ctx echo.Context) error {
	userRequest := response.CreateUserResponse{}
	if err := ctx.Bind(&userRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to create user",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusCreated, response.UserBaseResponse{
		Code:    http.StatusCreated,
		Message: "user created successfully",
		Data:    nil,
	})
}

func (h UserHandler) GetList(ctx echo.Context) error {
	users, err := h.userUsecase.GetList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list users failed",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no user found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got all users",
		Data:    users,
	})
}

// Get only one user by ID
func (h UserHandler) GetOne(ctx echo.Context) error {
	id := ctx.Param("id")
	users, err := h.userUsecase.GetOne(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to get specified user data",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no data found for specified user id",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got specified user data",
		Data:    users,
	})
}

// Update User
func (h UserHandler) UpdateData(ctx echo.Context) error {
	UpdateUser := response.UpdateUserResponse{}
	id := ctx.Param("id")

	if err := ctx.Bind(&UpdateUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "All required information must be defined",
			Data:    nil,
		})
	}

	if err := h.userUsecase.Update(id, UpdateUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to update data",
			Data:    nil,
		})
	}

	config.DB.Where("id = ?", id).Updates(&UpdateUser)
	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "Data updated succesfully",
		Data:    UpdateUser,
	})
}

// Delete User
// Get only one user by ID
func (h UserHandler) DeleteData(ctx echo.Context) error {
	id := ctx.Param("id")
	users, err := h.userUsecase.DeleteUser(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to delete data",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no data found for this id, no data deleted",
			Data:    nil,
		})
	}

	config.DB.Delete(&users, id)
	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully deleted user with the following data",
		Data:    users,
	})
}
