package postsmodel

import (
	"github.com/google/uuid"
)

type PostCreation struct {
	Images  []uuid.UUID `json:"images"`
	Content string      `json:"content"`
}
