package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/smochii/go-clean-architecture/config"
	"github.com/smochii/go-clean-architecture/domain"
	"github.com/smochii/go-clean-architecture/domain/value"
	"github.com/smochii/go-clean-architecture/logger"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	GetToken(email value.Email, password value.Password) (string, error)
}

type authService struct {
	userRepository domain.UserRepository
}

func (s *authService) GetToken(email value.Email, password value.Password) (string, error) {
	failAuthError := errors.New("Failed to auth")

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		logger.Debug(err.Error())
		return "", failAuthError
	}
	if user == nil {
		logger.Debug("User not found")
		return "", failAuthError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password().String()), []byte(password.String()))
	if err != nil {
		logger.Debug("Password is incorrect")
		return "", failAuthError
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"sub": user.Id().String(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenStr, err := token.SignedString([]byte(config.Conf.App.AuthSecret))
	if err != nil {
		return "", failAuthError
	}

	return tokenStr, nil
}

func NewAuthService(userRepository domain.UserRepository) IAuthService {
	return &authService{
		userRepository: userRepository,
	}
}
