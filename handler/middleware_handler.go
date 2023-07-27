package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/smochii/go-clean-architecture/config"
	"github.com/smochii/go-clean-architecture/domain/value"
	"github.com/smochii/go-clean-architecture/service"
)

type MiddlewareHandler struct {
	userService service.IUserService
}

func (h *MiddlewareHandler) AuthenticateMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Conf.App.AuthSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			sub := token.Claims.(jwt.MapClaims)["sub"].(string)
			userId := value.NewUserIdFromString(sub)
			user, err := h.userService.GetUser(userId)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Invalid token",
				})
			}
			c.Locals("me", user)
			return c.Next()
		},
	})
}

func NewMiddlewareHandler(userService service.IUserService) MiddlewareHandler {
	return MiddlewareHandler{
		userService: userService,
	}
}
