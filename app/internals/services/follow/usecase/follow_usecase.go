package followusecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
	"instagram/components/pubsub"
	"time"
)

type FollowUserUseCase interface {
	Execute(ctx context.Context, requester common.Requester, followingId uuid.UUID) (bool, error)
}

type followUserUseCase struct {
	followUserRepository FollowUserRepository
	pubsub               pubsub.MessageBroker
}

func NewFollowUserUseCase(followUserRepository FollowUserRepository, pubsub pubsub.MessageBroker) FollowUserUseCase {
	return &followUserUseCase{
		followUserRepository: followUserRepository,
		pubsub:               pubsub,
	}
}

type FollowUserRepository interface {
	CheckIsExist(ctx context.Context, uId, fId uuid.UUID) (bool, error)
	Create(ctx context.Context, dto *followusermodel.FollowUser) (bool, error)
	Delete(ctx context.Context, follower uuid.UUID, following uuid.UUID) (bool, error)
}

func (f followUserUseCase) Execute(ctx context.Context, requester common.Requester, followingId uuid.UUID) (bool, error) {
	exist, err := f.followUserRepository.CheckIsExist(ctx, requester.UserId(), followingId)

	if err != nil {
		return false, err
	}

	userId := requester.UserId()
	if userId == followingId {
		return false, common.NewCustomError(followusermodel.ErrCanFollowYourself, followusermodel.ErrCanFollowYourself.Error(), "can_follow_yourself")
	}

	if exist {
		// TODO: Implement unfollow here
		_, err := f.followUserRepository.Delete(ctx, userId, followingId)
		if err != nil {
			return false, err
		}

		// TODO: Public followed message here
		message := pubsub.NewAppMessage(&userId, common.UnfollowUserTopic, map[string]any{
			"follower_id":  userId,
			"following_id": followingId,
		})

		go func() {
			common.RecoverPanic(fmt.Sprintf("Error public message from topic %s", common.FollowedUserTopic))
			err := f.pubsub.Publish(ctx, message)
			if err != nil {
				fmt.Println(err)
			}
		}()

	} else {
		// TODO: Implement follow here
		now := time.Now().UTC()
		dto := &followusermodel.FollowUser{
			UserID:    userId,
			Following: followingId,
			CreatedAt: &now,
			UpdatedAt: &now,
		}

		_, err := f.followUserRepository.Create(ctx, dto)
		if err != nil {
			return false, err
		}

		// TODO: Public followed message here
		message := pubsub.NewAppMessage(&userId, common.FollowedUserTopic, map[string]any{
			"follower_id":  userId,
			"following_id": followingId,
		})

		go func() {
			common.RecoverPanic(fmt.Sprintf("Error public message from topic %s", common.FollowedUserTopic))
			err := f.pubsub.Publish(ctx, message)
			if err != nil {
				fmt.Println(err)
			}
		}()

	}

	return true, nil
}
