package usermodel

import "net/mail"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserLogin) Validate() error {
	if u.Email == "" {
		return UserEmailIsRequired
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return UserEmailInvalid
	}

	if u.Password == "" {
		return UserPasswordIsRequired
	}

	return nil
}
