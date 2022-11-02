package handler

import (
	"fiber/config"
	"fiber/entity/response"
	"fiber/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usercase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: usercase}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserResponse{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to create user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "user created successfully",
		Data:    nil,
	})
}

func (h UserHandler) GetList(ctx *fiber.Ctx) error {
	users, err := h.userUsecase.GetList()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "get list users failed",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "no user found",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "successfully got all users",
		Data:    users,
	})
}

// Get only one user by ID
func (h UserHandler) GetOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	users, err := h.userUsecase.GetOne(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to get specified user data",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "no data found for specified user id",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "successfully got specified user data",
		Data:    users,
	})
}

// Update User
func (h UserHandler) UpdateData(ctx *fiber.Ctx) error {
	UpdateUser := response.UpdateUserResponse{}
	id := ctx.Params("id")

	if err := ctx.BodyParser(&UpdateUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "All required information must be defined",
			Data:    nil,
		})
	}

	if err := h.userUsecase.Update(id, UpdateUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to update data",
			Data:    nil,
		})
	}

	config.DB.Where("id = ?", id).Updates(&UpdateUser)
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Data updated succesfully",
		Data:    UpdateUser,
	})
}

// Delete User
// Get only one user by ID
func (h UserHandler) DeleteData(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	users, err := h.userUsecase.DeleteUser(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to delete data",
			Data:    nil,
		})
	}

	if len(users) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "no data found for this id, no data deleted",
			Data:    nil,
		})
	}

	config.DB.Delete(&users, id)
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "successfully deleted user with the following data",
		Data:    users,
	})
}
