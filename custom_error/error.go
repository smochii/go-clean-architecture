package custom_error

import (
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
)

func NewInternalError(c fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": ErrorCodes.InternalError,
	})
}

func NewBadRequestError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": ErrorCodes.BadRequestError,
	})
}

func NewNotFoundError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": ErrorCodes.NotFoundError,
	})
}

func NewPermissionError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error": ErrorCodes.PermissionError,
	})
}

func NewUnauthorizedError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": ErrorCodes.UnauthorizedError,
	})
}

func NewDisabledAccountError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error": ErrorCodes.DisabledAccountError,
	})
}

func NewValidationError(c *fiber.Ctx, errors map[string]any) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": errors,
	})
}

func NewMaintenanceModeError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
		"error": ErrorCodes.MaintenanceModeError,
	})
}

func genValidationErrorData(errors validation.Errors) map[string]any {
	errorData := map[string]any{}
	for k, err := range errors {
		camelKey := toLowerCamelCase(k)
		switch e := err.(type) {
		case validation.Errors:
			errorData[camelKey] = genValidationErrorData(e)
		case validation.Error:
			errorData[camelKey] = e.Code()
		}
	}
	return errorData
}

func toLowerCamelCase(s string) string {
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
