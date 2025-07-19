package service

import (
	"fmt"

	"github.com/Humiditii/receept/internal/auth/model"
	"github.com/Humiditii/receept/internal/auth/repository"
	"github.com/Humiditii/receept/internal/email"
)

type authService struct {
	authRepo *repository.AuthRepository
	emailService *email.EmailService
}

type AuthService interface {
	Signup(user *model.User) (*model.User, error)
}

func NewAuthService (authRepo *repository.AuthRepository, emailService *email.EmailService) AuthService {
	return &authService{
		authRepo: authRepo,
		emailService: emailService,
	}
}

func (as *authService) Signup(user *model.User) (*model.User, error) {
	(*user).Verified = false
	result, err := (*as.authRepo).Create(user)

	if err != nil {
		return nil, err
	}

	subject := "Welcome to Receept!!!"
	body := fmt.Sprintf("Hey %s, welcome to our app, we are happy to have you onboard", user.Firstname)
	if err:= as.emailService.Send(user.Email,subject,body); err != nil{
		return nil, err
	}

	return result, nil
}