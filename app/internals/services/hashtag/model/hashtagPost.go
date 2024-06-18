
package hashtagmodel

import (
	"github.com/google/uuid"
	"time"
)

type HashtagPost struct {
	Hashtag_ID uuid.UUID `json:"hashtag_id"`
	Post_ID    uuid.UUID `json:"post_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}