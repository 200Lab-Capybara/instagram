package postusecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

type getListPostByUserIdUseCase struct {
	postRepository GetListPostRepository
	userRepo       GetUserByIdRepository
}

type GetListPostRepository interface {
	GetListPostByUserId(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]postsmodel.Post, error)
}

type GetUserByIdRepository interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*common.SimpleUser, error)
}

func NewGetListPostByUserIdUseCase(postRepository GetListPostRepository, userRepo GetUserByIdRepository) GetListPostByUserIdUseCase {
	return &getListPostByUserIdUseCase{
		postRepository: postRepository,
		userRepo:       userRepo,
	}
}

type GetListPostByUserIdUseCase interface {
	Execute(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]postsmodel.Post, error)
}

func (g getListPostByUserIdUseCase) Execute(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]postsmodel.Post, error) {
	simpleUser, err := g.userRepo.GetUserById(ctx, userId)

	if err != nil {
		return nil, err
	}

	fmt.Println(simpleUser.Status, "status:::::")

	if simpleUser.Status == common.UserDeleted.String() {
		fmt.Println("user deleted")
		return nil, common.NewCustomError(usermodel.ErrUserNotFound, usermodel.ErrUserNotFound.Error(), "user_deleted")
	}

	if simpleUser.Status == common.UserBanned.String() {
		return nil, common.NewCustomError(usermodel.ErrUserBanded, usermodel.ErrUserBanded.Error(), "user_banded")
	}

	paging.Process()

	data, err := g.postRepository.GetListPostByUserId(ctx, userId, paging)

	if err != nil {
		return nil, err
	}

	for i := range data {
		data[i].Owner = simpleUser
	}

	return data, nil
}
