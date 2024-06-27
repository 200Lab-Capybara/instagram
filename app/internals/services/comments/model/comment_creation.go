package modelcomment

import "github.com/google/uuid"

type CommentCreation struct {
	PostId  uuid.UUID
	Content string `json:"content"`
}
