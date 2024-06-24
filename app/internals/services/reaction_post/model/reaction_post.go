package reactionpostmodel

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrRecordReactPostNotFound = errors.New("record react story not found")
	ErrPostDoNotExists         = errors.New("post do not exists")
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
