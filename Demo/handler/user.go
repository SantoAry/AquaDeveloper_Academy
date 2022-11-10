// Handles response to client
package handler

import (
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

// Create a Role (Optional)
func (h UserHandler) CreateRole(ctx echo.Context) error {
	roleRequest := response.CreateRoleResponse{}
	if err := ctx.Bind(&roleRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.RoleBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.CreateRole(roleRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.RoleBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to create role",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusCreated, response.RoleBaseResponse{
		Code:    http.StatusCreated,
		Message: "role created successfully",
		Data:    nil,
	})
}

// Get list of all Roles
func (h UserHandler) GetAllRoles(ctx echo.Context) error {
	users, err := h.userUsecase.GetAllRoles()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list roles failed",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no role found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully got all roles",
		Data:    users,
	})
}

// Create new User
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

// Get list of all Users
func (h UserHandler) GetAllUsers(ctx echo.Context) error {
	role := ctx.QueryParam("role")
	users, err := h.userUsecase.GetAllUsers(role)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list users failed (find unsuccessful or role type invalid)",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no user found or role type invalid",
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
func (h UserHandler) GetOneUser(ctx echo.Context) error {
	id := ctx.Param("id")
	users, err := h.userUsecase.GetOneUser(id)
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

// Update User by ID
func (h UserHandler) UpdateUser(ctx echo.Context) error {
	UpdateUser := response.UpdateUserResponse{}
	id := ctx.Param("id")

	if err := ctx.Bind(&UpdateUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "All required information must be defined",
			Data:    nil,
		})
	}

	if err := h.userUsecase.UpdateUser(id, UpdateUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserBaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to update data or unauthorized access",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "Data updated succesfully",
		Data:    UpdateUser,
	})
}

// Delete user by ID
func (h UserHandler) DeleteUser(ctx echo.Context) error {
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

	return ctx.JSON(http.StatusOK, response.UserBaseResponse{
		Code:    http.StatusOK,
		Message: "successfully deleted user with the following data",
		Data:    users,
	})
}
