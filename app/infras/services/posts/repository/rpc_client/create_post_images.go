package postrpcclient

import (
	"context"
	"github.com/google/uuid"
	"instagram/common"
	"log"
)

type createPostImages struct {
	db common.SQLDatabase
}

func NewCreatePostImages(db common.SQLDatabase) *createPostImages {
	return &createPostImages{db: db}
}

func (c *createPostImages) CreatePostImages(ctx context.Context, postID uuid.UUID, images []uuid.UUID) error {
	// Do something
	log.Println("CreatePostImages")
	return nil
}
