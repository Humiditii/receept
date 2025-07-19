package service

import (
	"log"

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

// func (as *authService) Signup(user *model.User) (*model.User, error) {
// 	(*user).Verified = false
// 	result, err := (*as.authRepo).Create(user)

// 	if err != nil {
// 		return nil, err
// 	}

// 	welcomeMail := email.WelcomeMail(&user.Firstname)
// 	if err:= as.emailService.Send(user.Email,welcomeMail.Subject,welcomeMail.Body); err != nil{
// 		return nil, err
// 	}

// 	return result, nil
// }

func (as *authService) Signup(user *model.User) (*model.User, error) {
	user.Verified = false
	result, err := (*as.authRepo).Create(user)
	if err != nil {
		return nil, err
	}

	// Async email sending
	go func(to string, name string) {
		welcomeMail := email.WelcomeMail(&name)
		if err := as.emailService.Send(to, welcomeMail.Subject, welcomeMail.Body); err != nil {
			log.Printf("failed to send welcome email: %v\n", err)
		}
	}(user.Email, user.Firstname)

	return result, nil
}
