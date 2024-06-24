package usermodel

import (
	"errors"
	"github.com/google/uuid"
	"instagram/common"
	"time"
)

var (
	ErrUserNotFound            = errors.New("user not found")
	ErrUserAlreadyExists       = errors.New("user already exists")
	ErrUserEmailIsRequired     = errors.New("email is required")
	ErrUserEmailInvalid        = errors.New("email is invalid")
	ErrUserFirstNameIsRequired = errors.New("first name is required")
	ErrUserLastNameIsRequired  = errors.New("last name is required")
	ErrUserPasswordIsRequired  = errors.New("password is required")
	ErrUserPasswordLength      = errors.New("password must be at least 8 characters long")
	ErrInvalidEmailOrPass      = errors.New("invalid email or password")
	ErrUserBanded              = errors.New("user is baned")
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
