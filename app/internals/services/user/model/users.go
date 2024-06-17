package usermodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type UserRole int

const (
	RoleUser UserRole = iota + 1
	RoleModerator
	RoleAdmin
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	case RoleModerator:
		return "mod"
	default:
		return "user"
	}
}

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var r UserRole

	roleValue := string(bytes)

	if roleValue == "user" {
		r = RoleUser
	} else if roleValue == "admin" {
		r = RoleAdmin
	} else if roleValue == "mod" {
		r = RoleModerator
	}

	*role = r

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

// Add=1, Subtract=2, Multiply=3

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
	ID        uuid.UUID  `json:"id" gorm:"column:id"`
	Email     string     `json:"email" gorm:"column:email"`
	FirstName string     `json:"first_name" gorm:"column:first_name"`
	LastName  string     `json:"last_name" gorm:"column:last_name"`
	Password  string     `json:"password" gorm:"column:password"`
	Salt      string     `json:"salt" gorm:"column:salt"`
	Role      UserRole   `json:"role" gorm:"column:role"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
