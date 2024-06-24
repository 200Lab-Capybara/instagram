package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrStoryNotFound            = errors.New("story not found")
	ErrRecordReactStoryNotFound = errors.New("record react story not found")
)

type ReactionStory struct {
	UserId    uuid.UUID  `json:"user_id"`
	StoryId   uuid.UUID  `json:"story_id"`
	CreatedAt time.Time  `json:"create_at"`
	UpdatedAt *time.Time `json:"update_at"`
}

func (ReactionStory) TableName() string { return "react_stories" }

type Story struct {
	Id uuid.UUID `json:"id"`
}

func (Story) TableName() string {
	return "stories"
}
