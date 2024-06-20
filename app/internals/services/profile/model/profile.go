package model

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID             uuid.UUID      `json:"id" gorm:"colum:id"`
	DateOfBirth    time.Time      `json:"date_of_birth" gorm:"colum:date_of_birth"`
	UserId         uuid.UUID      `json:"userId" gorm:"colum:user_id"`
	Genders        *ProfileGender `json:"gender" gorm:"colum:gender"`
	Avatar         string         `json:"avatar" gorm:"colum:avatar"`
	CountFollowing int            `json:"count_following" gorm:"colum:count_following"`
	CountFollowers int            `json:"count_followers" gorm:"colum:count_followers"`
	CountPosts     int            `json:"count_posts" gorm:"colum:count_posts"`
	CreatedAt      time.Time      `json:"created_at" gorm:"colum:created_at"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"colum:updated_at"`
}

func (Profile) TableName() string {
	return "profiles"
}
