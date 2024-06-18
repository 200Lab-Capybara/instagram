package model

import "errors"

type ProfileGender int

const (
	Male ProfileGender = iota
	Female
	Other
)

var allGenders = [3]string{"Male", "Female", "Other"}

func (gender *ProfileGender) String() string {
	return allGenders[*gender]
}

func parseStr2ProfileGender(s string) (ProfileGender, error) {
	for i := range allGenders {
		if allGenders[i] == s {
			return ProfileGender(i), nil
		}
	}
	return ProfileGender(0), errors.New("invalid status string")
}
