package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smochii/go-clean-architecture/custom_error"
	"github.com/smochii/go-clean-architecture/handler/request"
	"github.com/smochii/go-clean-architecture/service"
)

type AuthHandler struct {
	authService service.IAuthService
}

func (h *AuthHandler) GenerateToken(c *fiber.Ctx) error {
	req := new(request.GetTokenRequest)
	if err := Bind(c, req); err != nil {
		return custom_error.NewBadRequestError(c)
	}

	if errors := Validate(req); errors != nil {
		return custom_error.NewValidationError(c, errors)
	}

	token, err := h.authService.GetToken(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	type LoginResponse struct {
		Token string `json:"token"`
	}
	return c.Status(fiber.StatusOK).JSON(LoginResponse{
		Token: token,
	})
}

func NewAuthHandler(authService service.IAuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}
