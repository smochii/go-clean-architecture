package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/smochii/go-clean-architecture/domain/value"
)

type GetTokenRequestBody struct {
	Email    value.Email
	Password value.Password
}

type GetTokenRequest struct {
	GetTokenRequestBody
}

func (r GetTokenRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email),
		validation.Field(&r.Password),
	)
}
