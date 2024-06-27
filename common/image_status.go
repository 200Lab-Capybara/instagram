package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ImageStatus int

const (
	ImageUsed ImageStatus = iota
	ImageUnUsed
)

func (status ImageStatus) String() string {
	switch status {
	case ImageUnUsed:
		return "unused"
	default:
		return "used"
	}
}
func (status *ImageStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var r ImageStatus

	statusValue := string(bytes)

	if statusValue == "used" {
		r = ImageUsed
	} else if statusValue == "unused" {
		r = ImageUnUsed
	}

	*status = r

	return nil
}
func (status *ImageStatus) Value() (driver.Value, error) {
	if status == nil {
		return nil, nil
	}
	return status.String(), nil
}
func (status *ImageStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", status.String())), nil
}
