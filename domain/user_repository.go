package domain

import (
	"github.com/smochii/go-clean-architecture/domain/entity"
	"github.com/smochii/go-clean-architecture/domain/value"
)

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	Update(user entity.User) (*entity.User, error)
	FindById(id value.UserId) (*entity.User, error)
	FindByEmail(email value.Email) (*entity.User, error)
	FindAll() ([]entity.User, error)
}
