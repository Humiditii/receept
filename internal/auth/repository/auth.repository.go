package repository

import (
	"github.com/Humiditii/receept/internal/auth/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	Create(user *model.User) (*model.User, error)
	FindOne(id int)  (*model.User, error)
}


func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (authR *authRepository) Create(user *model.User) (*model.User, error) {
	if err := authR.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func (authR *authRepository) FindOne(id int)  (*model.User, error) {

	var user model.User

	if err:=authR.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}