package service

import (
	"github.com/Humiditii/receept/internal/auth/model"
	"github.com/Humiditii/receept/internal/auth/repository"
)

type authService struct {
	authRepo *repository.AuthRepository
}

type AuthService interface {
	Signup(user *model.User) (*model.User, error)
}

func NewAuthService (authRepo *repository.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (as *authService) Signup(user *model.User) (*model.User, error) {
	(*user).Verified = false
	result, err := (*as.authRepo).Create(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}