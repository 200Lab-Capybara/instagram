package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type UserStatus int

const (
	Active UserStatus = iota + 1
	Inactive
	Banned
)

func (status UserStatus) String() string {
	switch status {
	case Banned:
		return "banned"
	case Inactive:
		return "inactive"
	default:
		return "active"
	}
}

func (status *UserStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var r UserStatus

	statusValue := string(bytes)

	if statusValue == "active" {
		r = Active
	} else if statusValue == "inactive" {
		r = Inactive
	} else if statusValue == "banned" {
		r = Banned
	}

	*status = r

	return nil
}

func (status *UserStatus) Value() (driver.Value, error) {
	if status == nil {
		return nil, nil
	}
	return status.String(), nil
}

func (status *UserStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", status.String())), nil
}
