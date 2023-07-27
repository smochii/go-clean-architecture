package service

import (
	"errors"
	"time"

	"github.com/smochii/go-clean-architecture/domain"
	"github.com/smochii/go-clean-architecture/domain/entity"
	"github.com/smochii/go-clean-architecture/domain/value"
	"github.com/smochii/go-clean-architecture/logger"
)

type IUserService interface {
	CreateUser(email value.Email, password value.Password) (*entity.User, error)
	UpdateUser(userId value.UserId, email value.Email, password value.Password) (*entity.User, error)
	GetUser(userId value.UserId) (*entity.User, error)
}

type userService struct {
	userRepository domain.UserRepository
}

func (r *userService) CreateUser(email value.Email, password value.Password) (*entity.User, error) {
	if r.existUser(email) {
		return nil, errors.New("Email is already registered")
	}

	hashedPassword, err := value.NewHashedPassword(password)
	if err != nil {
		logger.Debug(err.Error())
		return nil, errors.New("Failed to create user")
	}

	user := entity.NewUser(
		value.NewUserId(),
		email,
		hashedPassword,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		logger.Debug(err.Error())
		return nil, errors.New("Failed to create user")
	}
	created, err := r.userRepository.Create(user)
	if err != nil {
		logger.Debug(err.Error())
		return nil, errors.New("Failed to create user")
	}
	return created, nil
}

func (r *userService) UpdateUser(userId value.UserId, email value.Email, password value.Password) (*entity.User, error) {
	if r.existUser(email) {
		return nil, errors.New("Email is already registered")
	}

	hashedPassword, err := value.NewHashedPassword(password)
	if err != nil {
		logger.Debug(err.Error())
		return nil, errors.New("Failed to update user")
	}

	user, err := r.userRepository.FindById(userId)
	if err != nil {
		logger.Debug(err.Error())
		return nil, errors.New("Failed to update user")
	}
	if user == nil {
		return nil, errors.New("User not found")
	}

	user.SetEmail(email).SetPassword(hashedPassword).SetUpdatedAt(time.Now())

	updated, err := r.userRepository.Update(*user)
	if err != nil {
		logger.Debug(err.Error())
		return nil, errors.New("Failed to update user")
	}
	return updated, nil
}

func (r *userService) GetUser(userId value.UserId) (*entity.User, error) {
	user, err := r.userRepository.FindById(userId)
	if err != nil {
		logger.Debug(err.Error())
		return nil, errors.New("Failed to get user")
	}
	if user == nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (r *userService) existUser(email value.Email) bool {
	user, err := r.userRepository.FindByEmail(email)
	if err != nil {
		logger.Debug(err.Error())
		return false
	}
	if user == nil {
		return false
	}
	return true
}

func NewUserService(userRepository domain.UserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}
