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

type Story struct {
	Id          uuid.UUID    `json:"id" gorm:"id"`
	UserId      uuid.UUID    `json:"user_id" gorm:"user_id"`
	Content     string       `json:"content_story" gorm:"content_story"`
	ReactCount  int          `json:"react_count" gorm:"react_count"`
	ExpiresTime int          `json:"expires_time" gorm:"expires_time"`
	Image       common.Image `json:"image" gorm:"image"`
	CreatedAt   time.Time    `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"updated_at"`
	IsActive    bool         `json:"is_active" gorm:"is_active"`
}

func (Story) TableName() string { return "stories" }
