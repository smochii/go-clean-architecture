package handler

import (
	"errors"
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/smochii/go-clean-architecture/logger"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func Bind(c *fiber.Ctx, req interface{}) error {
	if err := c.ParamsParser(req); err != nil {
		logger.Debug(err.Error())
		return errors.New("Failed to parse request params")
	}

	if err := c.QueryParser(req); err != nil {
		logger.Debug(err.Error())
		return errors.New("Failed to parse request query")
	}

	if err := c.BodyParser(req); err != nil {
		logger.Debug(err.Error())
		return errors.New("Failed to parse request body")
	}

	return nil
}

func Validate(req validation.Validatable) map[string]any {
	err := req.Validate()

	var validationErrors validation.Errors
	if errors.As(err, &validationErrors) {
		errorData := genValidationErrorData(validationErrors)
		return errorData
	}

	return nil
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
