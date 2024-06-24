package usermodel

import (
	"net/mail"
)

type UserCreation struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (u *UserCreation) Validate() error {
	if u.Email == "" {
		return ErrUserEmailIsRequired
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return ErrUserEmailInvalid
	}

	if u.FirstName == "" {
		return ErrUserFirstNameIsRequired
	}

	if u.LastName == "" {
		return ErrUserLastNameIsRequired
	}

	if u.Password == "" {
		return ErrUserPasswordIsRequired
	}

	if len(u.Password) < 8 {
		return ErrUserPasswordLength
	}

	return nil
}
