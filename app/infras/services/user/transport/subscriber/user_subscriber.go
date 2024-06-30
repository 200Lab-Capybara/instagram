package usersubscriber

import (
	"context"
	"fmt"
	usermodel "instagram/app/internals/services/user/model"
	userusecase "instagram/app/internals/services/user/usecase"
	"instagram/common"
	"instagram/components/pubsub"
)

type postSubscriber struct {
	pubsub          pubsub.MessageBroker
	followUseCase   userusecase.UpdateFollowUseCase
	unfollowUseCase userusecase.UpdateUnfollowUseCase
}

func NewUserSubscriber(pubsub pubsub.MessageBroker, unfollowUseCase userusecase.UpdateUnfollowUseCase, followUseCase userusecase.UpdateFollowUseCase) *postSubscriber {
	return &postSubscriber{
		pubsub:          pubsub,
		unfollowUseCase: unfollowUseCase,
		followUseCase:   followUseCase,
	}
}

func (sbr *postSubscriber) Init() {
	go func() {
		ctx := context.Background()
		subscribe, c, err := sbr.pubsub.Subscribe(ctx, common.FollowedUserTopic)
		if err != nil {
			return
		}

		defer c()

		for msg := range subscribe {
			payload := msg.Payload
			dto := &usermodel.FollowUserPayload{
				UserID:      payload["follower_id"].(string),
				FollowingID: payload["following_id"].(string),
			}

			_, err := sbr.followUseCase.Execute(ctx, dto)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	go func() {
		ctx := context.Background()
		subscribe, c, err := sbr.pubsub.Subscribe(ctx, common.UnfollowUserTopic)
		if err != nil {
			return
		}

		defer c()

		for msg := range subscribe {
			payload := msg.Payload

			dto := &usermodel.FollowUserPayload{
				UserID:      payload["follower_id"].(string),
				FollowingID: payload["following_id"].(string),
			}

			_, err := sbr.unfollowUseCase.Execute(ctx, dto)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}
