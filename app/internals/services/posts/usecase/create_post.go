package postusecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
	"instagram/components/pubsub"
	"time"
)

type createPostUseCase struct {
	postRepository CreatePostRepository
	pubsub         pubsub.MessageBroker
}

func NewCreatePostUseCase(postRepository CreatePostRepository, pubsub pubsub.MessageBroker) CreatePostUseCase {
	return &createPostUseCase{
		postRepository: postRepository,
		pubsub:         pubsub,
	}
}

type CreatePostRepository interface {
	CreatePost(ctx context.Context, post *postsmodel.Post) (*uuid.UUID, error)
}

type CreatePostUseCase interface {
	Execute(ctx context.Context, requester common.Requester, dto *postsmodel.PostCreation) (*uuid.UUID, error)
}

func (c *createPostUseCase) Execute(ctx context.Context, requester common.Requester, dto *postsmodel.PostCreation) (*uuid.UUID, error) {
	postID, _ := uuid.NewV7()
	userID := requester.UserId()
	usedHashtag := false

	hashtags := common.GetHashtag(dto.Content)
	if len(hashtags) > 0 {
		usedHashtag = true
	}

	post := &postsmodel.Post{
		ID:           postID,
		UserID:       userID,
		Content:      dto.Content,
		LikeCount:    0,
		CommentCount: 0,
		Status:       postsmodel.PostActive,
		UsedHashtag:  usedHashtag,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	id, err := c.postRepository.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	go func() {
		common.RecoverPanic(fmt.Sprintf("Error public message from topic %s", common.CreatedPostTopic))
		postMessage := pubsub.NewAppMessage(&userID, common.CreatedPostTopic, map[string]interface{}{
			"post_id":  postID,
			"user_id":  userID,
			"hashtags": hashtags,
			"images":   dto.Images,
		})

		// TODO: Publish CreatedPostTopic event
		err := c.pubsub.Publish(ctx, postMessage)
		if err != nil {
			panic(err)
		}
	}()

	return id, nil
}
