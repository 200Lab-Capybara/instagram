package postusecase

import (
	"context"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
	"time"
)

type createPostUseCase struct {
	postRepository CreatePostRepository
	pImagesRepo    CreatePostImagesRepository
}

func NewCreatePostUseCase(postRepository CreatePostRepository, pImagesRepo CreatePostImagesRepository) CreatePostUseCase {
	return &createPostUseCase{
		postRepository: postRepository,
		pImagesRepo:    pImagesRepo,
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

func (c createPostUseCase) Execute(ctx context.Context, requester common.Requester, dto *postsmodel.PostCreation) (*uuid.UUID, error) {
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
		return nil, nil
	}

	// TODO: Publish event

	return id, nil
}
