package reactionpostusecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	"instagram/common"
	"instagram/components/pubsub"
)

type reactionPostUseCase struct {
	reactionPostRepository ReactionPostRepository
	postRepository         GetPostRepository
	pubsub                 pubsub.MessageBroker
}

func NewLikePostUseCase(reactionPostRepository ReactionPostRepository, postRepository GetPostRepository, pubsub pubsub.MessageBroker) ReactionPostUseCase {
	return &reactionPostUseCase{
		reactionPostRepository: reactionPostRepository,
		postRepository:         postRepository,
		pubsub:                 pubsub,
	}
}

func (uc *reactionPostUseCase) Execute(ctx context.Context, requester common.Requester, postId uuid.UUID) (bool, error) {
	post, err := uc.postRepository.FindById(ctx, postId)
	userId := requester.UserId()
	reactType := common.ReactedPostLike

	if err != nil {
		return false, err
	}

	if post.Status == "deleted" {
		//TODO
	}
	checkExist, err := uc.reactionPostRepository.CheckExistReactionPost(ctx, userId, postId)

	fmt.Println(checkExist, "checkExist")
	if checkExist {
		_, err = uc.reactionPostRepository.RemoveRecordReactionPost(ctx, postId, userId)
		reactType = common.ReactedPostUnlike
		if err != nil {
			return false, err
		}
		//increase count reaction post
		//TODO
	} else {
		_, err = uc.reactionPostRepository.CreateNewReactionPost(ctx, userId, postId)
		if err != nil {
			return false, err
		}
	}

	fmt.Println(checkExist)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Error public message from topic %s", common.ReactedPostTopic)
			}
		}()

		postMessage := pubsub.NewAppMessage(&userId, common.ReactedPostTopic, map[string]interface{}{
			"post_id":    postId,
			"react_type": reactType,
		})

		// TODO: Publish CreatedPostTopic event
		err := uc.pubsub.Publish(ctx, postMessage)
		if err != nil {
			panic(err)
		}
	}()

	return true, nil
}

type ReactionPostUseCase interface {
	Execute(ctx context.Context, requester common.Requester, postId uuid.UUID) (bool, error)
}

type ReactionPostRepository interface {
	CreateNewReactionPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error)
	RemoveRecordReactionPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error)
	CheckExistReactionPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error)
}

type GetPostRepository interface {
	FindById(ctx context.Context, postId uuid.UUID) (*reactionpostmodel.Post, error)
}
