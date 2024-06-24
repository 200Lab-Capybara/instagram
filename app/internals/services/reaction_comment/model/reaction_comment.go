package modelreactioncomment

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrUserNotFound               = errors.New("user not found")
	ErrCommentNotFound            = errors.New("comment not found")
	ErrRecordReactCommentNotFound = errors.New("record react comment not found")
)

type ReactionComment struct {
	UserId    uuid.UUID  `json:"user_id"`
	CommentId uuid.UUID  `json:"comment_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (ReactionComment) TableName() string {
	return "reaction_comments"
}

type Comment struct {
	Id uuid.UUID `json:"id"`
}

func (Comment) TableName() string {
	return "comments"
}
