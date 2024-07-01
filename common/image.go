package common

import (
	"database/sql/driver"
	"encoding/json"
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

func (status ImageStatus) MarshalJSON() ([]byte, error) {
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
	Id          uuid.UUID   `json:"id"`
	UserId      uuid.UUID   `json:"user_id"`
	ImageUrl    string      `json:"image_url"`
	Size        int         `json:"size"`
	Width       int         `json:"width"`
	Height      int         `json:"height"`
	Status      ImageStatus `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	StorageName string      `json:"storage_name"` // S3
	Extension   string      `json:"extension"`    // jpg, png, ...
}

func (image *Image) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), image)
}

func (image Image) Value() (driver.Value, error) {
	return json.Marshal(image)
}
func (Image) TableName() string {
	return "images"
}
