package usermodel

import (
	"errors"
	"github.com/google/uuid"
	"instagram/common"
	"time"
)

var (
	UserNotFound            = errors.New("user not found")
	UserAlreadyExists       = errors.New("user already exists")
	UserEmailIsRequired     = errors.New("email is required")
	UserEmailInvalid        = errors.New("email is invalid")
	UserFirstNameIsRequired = errors.New("first name is required")
	UserLastNameIsRequired  = errors.New("last name is required")
	UserPasswordIsRequired  = errors.New("password is required")
	UserPasswordLength      = errors.New("password must be at least 8 characters long")
	ErrorInvalidEmailOrPass = errors.New("invalid email or password")
	ErrorUserBanded         = errors.New("user is baned")
)

type User struct {
	ID        uuid.UUID         `json:"id" gorm:"column:id"`
	Email     string            `json:"email" gorm:"column:email"`
	FirstName string            `json:"first_name" gorm:"column:first_name"`
	LastName  string            `json:"last_name" gorm:"column:last_name"`
	Password  string            `json:"password" gorm:"column:password"`
	Salt      string            `json:"salt" gorm:"column:salt"`
	Role      common.UserRole   `json:"role" gorm:"column:role"`
	Status    common.UserStatus `json:"status" gorm:"column:status"`
	CreatedAt *time.Time        `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time        `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
