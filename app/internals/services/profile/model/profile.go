package model

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID             uuid.UUID      `json:"id" gorm:"column:id"`
	DateOfBirth    time.Time      `json:"date_of_birth" gorm:"column:date_of_birth"`
	UserId         uuid.UUID      `json:"userId" gorm:"column:user_id"`
	Gender         *ProfileGender `json:"gender" gorm:"column:gender"`
	Avatar         string         `json:"avatar" gorm:"column:avatar"`
	CountFollowing int            `json:"count_following" gorm:"column:count_following"`
	CountFollowers int            `json:"count_followers" gorm:"column:count_followers"`
	CountPosts     int            `json:"count_posts" gorm:"column:count_posts"`
	CreatedAt      time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"column:updated_at"`
}

func (Profile) TableName() string {
	return "profiles"
}
