package postusecase

import (
	"context"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

type getFeedUseCase struct {
	userRepo   GetUserRepository
	postRepo   GetPostRepository
	followRepo GetFollowingRepository
}

type GetFeedUseCase interface {
	Execute(ctx context.Context, requester common.Requester, paging *common.Paging) ([]postsmodel.Post, error)
}

type GetPostRepository interface {
	GetPostsByUserIds(ctx context.Context, ids []uuid.UUID, paging *common.Paging) ([]postsmodel.Post, error)
}

type GetFollowingRepository interface {
	GetListFollowingByUserId(ctx context.Context, uid uuid.UUID) ([]uuid.UUID, error)
}

type GetUserRepository interface {
	GetUserByIds(ctx context.Context, ids []uuid.UUID) ([]common.SimpleUser, error)
}

func NewGetFeedUseCase(userRepo GetUserRepository, postRepo GetPostRepository, followRepo GetFollowingRepository) GetFeedUseCase {
	return &getFeedUseCase{
		userRepo:   userRepo,
		postRepo:   postRepo,
		followRepo: followRepo,
	}
}

func (g getFeedUseCase) Execute(ctx context.Context, requester common.Requester, paging *common.Paging) ([]postsmodel.Post, error) {
	// Get user's following
	paging.Process()
	following, err := g.followRepo.GetListFollowingByUserId(ctx, requester.UserId())
	if err != nil {
		return nil, err
	}

	// Get posts from following
	posts, err := g.postRepo.GetPostsByUserIds(ctx, following, paging)

	if err != nil {
		return nil, err
	}

	// Get user's posts
	setIds := make(map[uuid.UUID]bool)

	for _, post := range posts {
		setIds[post.UserID] = true
	}

	i := 0
	listIds := make([]uuid.UUID, len(setIds))
	for id := range setIds {
		listIds[i] = id
		i++
	}

	// Get user's info
	users, err := g.userRepo.GetUserByIds(ctx, listIds)
	if err != nil {
		return nil, err
	}

	userHashtable := make(map[uuid.UUID]common.SimpleUser)
	for _, user := range users {
		userHashtable[user.UserId] = user
	}

	for i, post := range posts {
		simpleUser := userHashtable[post.UserID]
		posts[i].Owner = &simpleUser
	}

	// Combine and sort
	// Return the result
	return posts, nil
}
