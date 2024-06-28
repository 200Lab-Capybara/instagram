package modelcomment

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrPostNotFound = errors.New("post not found")
)

type Comment struct {
	Id         uuid.UUID `json:"id" gorm:"column:id"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
	Content    string    `json:"content" gorm:"column:content"`
	UserId     uuid.UUID `json:"user_id" gorm:"column:user_id"`
	PostId     uuid.UUID `json:"post_id" gorm:"column:post_id"`
	ReactCount int       `json:"react_count" gorm:"column:react_count"`
}

type Post struct {
	Id uuid.UUID `json:"post_id"`
}

func (Comment) TableName() string { return "comments" }

func (Post) TableName() string { return "posts" }
