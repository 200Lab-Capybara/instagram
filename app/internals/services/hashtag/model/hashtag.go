
package hashtagmodel

import (
	"github.com/google/uuid"
	"time"
)

type Hashtag struct {
	ID        uuid.UUID  `json:"id"`
	Hashtag   string     `json:"hashtag"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
