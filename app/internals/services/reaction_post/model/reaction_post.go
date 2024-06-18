package model

import (
	"github.com/google/uuid"
	"time"
)

type ReactionPost struct {
	PostID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
}

type Post struct {
	ID     uuid.UUID
	Status string
}

func (ReactionPost) TableName() string {
	return "reaction_posts"
}
