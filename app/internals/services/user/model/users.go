package usermodel

import (
	"errors"
	"github.com/google/uuid"
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
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Password  string     `json:"password"`
	Salt      string     `json:"salt"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
