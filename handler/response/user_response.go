package response

import (
	"time"

	"github.com/smochii/go-clean-architecture/domain/value"
)

type UserResponseSchema struct {
	Id        value.UserId `json:"id"`
	Email     value.Email  `json:"email"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

type CreateUserResponse UserResponseSchema
type UpdateUserResponse UserResponseSchema
type GetUserResponse UserResponseSchema
type GetMeResponse UserResponseSchema
