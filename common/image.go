package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
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

func (status *ImageStatus) UnmarshalJSON(b []byte) error {
	statusValue := string(b)
	if statusValue == "\"used\"" {
		*status = ImageUsed
	} else if statusValue == "\"unused\"" {
		*status = ImageUnUsed
	}

	return nil
}

type Image struct {
	Id          uuid.UUID   `json:"id" gorm:"id"`
	UserId      uuid.UUID   `json:"user_id" gorm:"user_id"`
	ImageUrl    string      `json:"image_url" gorm:"image_url"`
	Size        int         `json:"size" gorm:"size"`
	Width       int         `json:"width" gorm:"width"`
	Height      int         `json:"height" gorm:"height"`
	Status      ImageStatus `json:"status" gorm:"status"`
	CreatedAt   time.Time   `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"updated_at"`
	StorageName string      `json:"storage_name" gorm:"storage_name"` // S3
	Extension   string      `json:"extension" gorm:"extension"`       // jpg, png, ...
}
