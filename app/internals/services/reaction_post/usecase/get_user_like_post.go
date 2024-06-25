package reactionpostusecase


type getUserLikePostUseCase struct{
	getUserLikePostRepo GetUserLikePostRepo
	postRepository 		GetPostRepository
}

func GetUserLikePostUC(getUserLikePostRepo GetUserLikePostRepo, postRepository GetPostRepository){
	return &getUserLikePostUseCase{
		getUserLikePostRepo: 	getUserLikePostRepo
		postRepository: 		postRepository
	}
}

func (uc *getUserLikePostUseCase) Execute(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (any, error){
	post, err := uc.postRepository.FindById(ctx, post_id)
	if err!=nil{
		return false, err
	}

	if post.Status == "deleted" {
		return false, common.ErrInvalidRequest(reactionpostmodel.ErrPostDoNotExist)
	}

	listUser, err := ListingUserLikePost(ctx, userId, postId)
	if err != nil {
		return false, err
	}

	return listUser, nil

}


type GetUserLikePostUseCase interface{
	Execute(ctx context.Context, userId uuid.UUID, post_id uuid.UUID) (any, error)
} 

type GetUserLikePostRepo interface {
	ListingUserLikePost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (any, error)
}