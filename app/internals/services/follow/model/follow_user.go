package followusermodel

import (
	"errors"
	"github.com/google/uuid"
	"instagram/common"
	"time"
)

var (
	ErrInvalidUserId      = errors.New("invalid user id")
	ErrInvalidFollowingId = errors.New("invalid following id")
	ErrCanFollowYourself  = errors.New("you can't follow yourself")
)

type FollowUser struct {
	UserID    uuid.UUID  `json:"user_id" gorm:"column:user_id"`
	Following uuid.UUID  `json:"following" gorm:"column:following"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type Follower struct {
	common.SimpleUser
	Followed bool `json:"followed"`
}

func (FollowUser) TableName() string {
	return "follow_users"
}
