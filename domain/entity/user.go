package entity

import (
	"time"

	"github.com/smochii/go-clean-architecture/domain/value"
)

type User struct {
	id        value.UserId
	email     value.Email
	password  value.HashedPassword
	createdAt time.Time
	updatedAt time.Time
}

func (u *User) Id() value.UserId {
	return u.id
}

func (u *User) Email() value.Email {
	return u.email
}

func (u *User) Password() value.HashedPassword {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetEmail(email value.Email) *User {
	u.email = email
	return u
}

func (u *User) SetPassword(password value.HashedPassword) *User {
	u.password = password
	return u
}

func (u *User) SetUpdatedAt(updatedAt time.Time) *User {
	u.updatedAt = updatedAt
	return u
}

func NewUser(id value.UserId, email value.Email, password value.HashedPassword, createdAt time.Time, updatedAt time.Time) User {
	return User{
		id:        id,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
