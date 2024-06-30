package postsmodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"instagram/common"
	"time"
)

type PostStatus int

var (
	ErrPostNotFound = errors.New("post not found")
)

const (
	PostActive PostStatus = iota + 1
	PostDeleted
)

func (status PostStatus) String() string {
	switch status {
	case PostDeleted:
		return "deleted"
	default:
		return "active"
	}
}

func (status *PostStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var r PostStatus

	statusValue := string(bytes)

	if statusValue == "active" {
		r = PostActive
	} else if statusValue == "deleted" {
		r = PostDeleted
	}

	*status = r

	return nil
}

func (status *PostStatus) Value() (driver.Value, error) {
	if status == nil {
		return nil, nil
	}
	return status.String(), nil
}

func (status *PostStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", status.String())), nil
}

type Post struct {
	ID           uuid.UUID          `json:"id" gorm:"column:id"`
	UserID       uuid.UUID          `json:"-" gorm:"column:user_id"`
	Owner        *common.SimpleUser `json:"owner" gorm:"-"`
	Content      string             `json:"content" gorm:"column:content"`
	LikeCount    int                `json:"like_count" gorm:"column:like_count"`
	CommentCount int                `json:"comment_count" gorm:"column:comment_count"`
	Status       PostStatus         `json:"-" gorm:"column:status"`
	UsedHashtag  bool               `json:"-" gorm:"column:used_hashtag"`
	CreatedAt    time.Time          `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time          `json:"updated_at" gorm:"column:updated_at"`
}

func (Post) TableName() string {
	return "posts"
}
