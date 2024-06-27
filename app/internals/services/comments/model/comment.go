package modelcomment

import "time"

type Comment struct {
	Id         string    `json:"id" gorm:"column:id"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
	Content    string    `json:"content" gorm:"column:content"`
	UserId     string    `json:"user_id" gorm:"column:user_id"`
	PostId     string    `json:"post_id" gorm:"column:post_id"`
	ReactCount int       `json:"react_count" gorm:"column:react_count"`
}

func (Comment) TableName() string { return "comments" }
