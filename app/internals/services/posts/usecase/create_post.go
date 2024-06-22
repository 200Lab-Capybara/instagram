package postusecase

import (
	"context"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
	"instagram/components/pubsub"
	"time"
)

type createPostUseCase struct {
	postRepository CreatePostRepository
	pImagesRepo    CreatePostImagesRepository
	pubsub         pubsub.MessageBroker
}

func NewCreatePostUseCase(postRepository CreatePostRepository, pImagesRepo CreatePostImagesRepository, pubsub pubsub.MessageBroker) CreatePostUseCase {
	return &createPostUseCase{
		postRepository: postRepository,
		pImagesRepo:    pImagesRepo,
		pubsub:         pubsub,
	}
}

type CreatePostRepository interface {
	CreatePost(ctx context.Context, post *postsmodel.Post) (*uuid.UUID, error)
}

type CreatePostImagesRepository interface {
	CreatePostImages(ctx context.Context, postID uuid.UUID, images []uuid.UUID) error
}

type CreatePostUseCase interface {
	Execute(ctx context.Context, requester common.Requester, dto *postsmodel.PostCreation) (*uuid.UUID, error)
}

func (c *createPostUseCase) Execute(ctx context.Context, requester common.Requester, dto *postsmodel.PostCreation) (*uuid.UUID, error) {
	postID, _ := uuid.NewV7()
	userID := requester.UserId()

	if len(dto.Images) > 0 {
		// Create post images
		err := c.pImagesRepo.CreatePostImages(ctx, postID, dto.Images)
		if err != nil {
			return nil, err
		}
	}

	post := &postsmodel.Post{
		ID:           postID,
		UserID:       userID,
		Content:      dto.Content,
		LikeCount:    0,
		CommentCount: 0,
		Status:       postsmodel.PostActive,
		UsedHashtag:  false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	id, err := c.postRepository.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	postMessage := pubsub.NewAppMessage(&userID, common.CreatedPostTopic, map[string]interface{}{
		"post_id": postID,
		"user_id": userID,
	})

	// TODO: Publish CreatedPostTopic event
	err = c.pubsub.Publish(ctx, postMessage)
	if err != nil {
		return nil, err
	}

	return id, nil
}
