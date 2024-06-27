package model

import (
	"errors"
	"github.com/google/uuid"
	"instagram/common"
	"time"
)

var (
	ErrStoryNotFound = errors.New("story not found")
)

type ImageStory struct {
	Id          uuid.UUID          `json:"id"`
	UserId      uuid.UUID          `json:"user_id"`
	ImageUrl    string             `json:"image_url"`
	Size        int                `json:"size"`
	Width       int                `json:"width"`
	Height      int                `json:"height"`
	Status      common.ImageStatus `json:"status"`
	CreatedAt   time.Time          `json:"create_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	StorageName string             `json:"storage_name"`
}
type Story struct {
	Id          uuid.UUID  `json:"id" gorm:"id"`
	UserId      uuid.UUID  `json:"user_id" gorm:"user_id"`
	Content     string     `json:"content_story" gorm:"content_story"`
	Count       int        `json:"react_count" gorm:"react_count"`
	ExpiresTime int        `json:"expires_time" gorm:"expires_time"`
	Image       ImageStory `json:"image" gorm:"image"`
	CreatedAt   time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"updated_at"`
	IsActive    bool       `json:"is_active" gorm:"is_active"`
}

func (Story) TableName() string { return "stories" }
