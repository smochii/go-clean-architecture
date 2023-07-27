package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/smochii/go-clean-architecture/custom_error"
	"github.com/smochii/go-clean-architecture/domain/entity"
	"github.com/smochii/go-clean-architecture/handler/request"
	"github.com/smochii/go-clean-architecture/handler/response"
	"github.com/smochii/go-clean-architecture/logger"
	"github.com/smochii/go-clean-architecture/service"
)

type UserHandler struct {
	userService service.IUserService
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	req := new(request.CreateUserRequest)
	if err := Bind(c, req); err != nil {
		return custom_error.NewBadRequestError(c)
	}

	if errors := Validate(req); errors != nil {
		return custom_error.NewValidationError(c, errors)
	}

	user, err := h.userService.CreateUser(req.Email, req.Password)
	if err != nil {
		logger.Debug(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.CreateUserResponse{
		Id:        user.Id(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	})
}

func (h *UserHandler) Me(c *fiber.Ctx) error {
	user := c.Locals("me").(*entity.User)

	return c.Status(fiber.StatusOK).JSON(response.GetMeResponse{
		Id:        user.Id(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	})
}

func NewUserHandler(userService service.IUserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}
