package model

import "time"

type UserType string


const (
	Admin UserType = "admin"
	Seller  UserType = "seller"
	Buyer UserType = "buyer"
)

type User struct {
	ID          uint      ` json:"id" gorm:"primaryKey"`
	Firstname   string    `json:"firstname" gorm:"not null"`
	Lastname    string    `json:"lastname" gorm:"not null"`
	Email       string    `json:"email" gorm:"uniqueIndex;not null"`
	Phone       string    `json:"phone" gorm:"uniqueIndex;not null"`
	AccountType UserType  `json:"account_type" gorm:"type:varchar(20);not null"`
	Verified    bool      `json:"verified" gorm:"default:false"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}