package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	UserNotFound  = errors.New("user not found")
	StoryNotFound = errors.New("story not found")
)

type ReactionStory struct {
	UserId     uuid.UUID  `json:"user_id"`
	StoryId    uuid.UUID  `json:"story_id"`
	Created_At time.Time  `json:"create_at"`
	Updated_At *time.Time `json:"update_at"`
}

func (ReactionStory) TableName() string { return "react_story" }

type Story struct {
	Id uuid.UUID `json:"id"`
}
