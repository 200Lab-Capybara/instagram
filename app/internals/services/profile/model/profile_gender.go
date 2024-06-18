package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

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

func (gender *ProfileGender) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	v, err := parseStr2ProfileGender(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	*gender = v

	return nil
}

func (gender *ProfileGender) Value() (driver.Value, error) {
	if gender == nil {
		return nil, nil
	}
	return gender.String(), nil
}

func (gender *ProfileGender) MarshalJSON() ([]byte, error) {
	if gender == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", gender.String())), nil
}

func (gender *ProfileGender) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	genderValue, err := parseStr2ProfileGender(str)

	if err != nil {
		return err
	}

	*gender = genderValue
	return nil
}
