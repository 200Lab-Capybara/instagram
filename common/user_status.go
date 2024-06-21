package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type UserStatus int

const (
	UserActive UserStatus = iota + 1
	UserInactive
	UserBanned
	UserDeleted
)

func (status UserStatus) String() string {
	switch status {
	case UserBanned:
		return "banned"
	case UserInactive:
		return "inactive"
	case UserDeleted:
		return "deleted"
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
		r = UserActive
	} else if statusValue == "inactive" {
		r = UserInactive
	} else if statusValue == "banned" {
		r = UserBanned
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
