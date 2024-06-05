package common

import (
	"github.com/google/uuid"
	"time"
)

type SQLModel struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewSQLModel() SQLModel {
	id, _ := uuid.NewV7()
	now := time.Now().UTC()

	return SQLModel{
		ID:        id,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
