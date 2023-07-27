package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/smochii/go-clean-architecture/domain/value"
)

type CreateUserRequestBody struct {
	Email    value.Email
	Password value.Password
}

type CreateUserRequest struct {
	CreateUserRequestBody
}

func (r CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email),
		validation.Field(&r.Password),
	)
}

type UpdateUserRequestParam struct {
	UserId value.UserId
}

type UpdateUserRequestBody struct {
	Email    value.Email
	Password value.Password
}

type UpdateUserRequest struct {
	UpdateUserRequestParam
	UpdateUserRequestBody
}

func (r UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email),
		validation.Field(&r.Password),
	)
}

type GetUserRequestParam struct {
	UserId value.UserId `params:"userId"`
}

type GetUserRequest struct {
	GetUserRequestParam
}

func (r GetUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserId),
	)
}
