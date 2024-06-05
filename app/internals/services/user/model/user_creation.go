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
		return UserEmailIsRequired
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return UserEmailInvalid
	}

	if u.FirstName == "" {
		return UserFirstNameIsRequired
	}

	if u.LastName == "" {
		return UserLastNameIsRequired
	}

	if u.Password == "" {
		return UserPasswordIsRequired
	}

	if len(u.Password) < 8 {
		return UserPasswordLength
	}

	return nil
}
