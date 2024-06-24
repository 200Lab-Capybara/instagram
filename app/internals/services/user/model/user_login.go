package usermodel

import "net/mail"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserLogin) Validate() error {
	if u.Email == "" {
		return ErrUserEmailIsRequired
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return ErrUserEmailInvalid
	}

	if u.Password == "" {
		return ErrUserPasswordIsRequired
	}

	return nil
}
