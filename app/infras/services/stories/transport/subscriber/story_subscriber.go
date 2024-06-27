package storysubscriber

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/stories/usecase"
	"instagram/common"
	"instagram/components/pubsub"
	"log"
)

type storySubscriber struct {
	pubsub       pubsub.MessageBroker
	increaselike usecase.IncreaseLikeCountUseCase
	decreaselike usecase.DecreaseLikeCountUseCase
}

func NewStorySubscriber(ps pubsub.MessageBroker, increase usecase.IncreaseLikeCountUseCase, decrease usecase.DecreaseLikeCountUseCase) *storySubscriber {
	return &storySubscriber{
		pubsub:       ps,
		increaselike: increase,
		decreaselike: decrease,
	}

}
func (sbr *storySubscriber) Init() {
	go func() {
		ctx := context.Background()
		subscribe, c, err := sbr.pubsub.Subscribe(ctx, common.ReactedStoryTopic)
		if err != nil {
			return
		}

		defer c()

		for msg := range subscribe {
			if msg == nil {
				return
			}

			payload := msg.Payload
			StoryId := uuid.MustParse(payload["story_id"].(string))
			reactType := payload["react_type"]

			switch reactType {
			case common.ReactedStoryLike:
				_, err := sbr.increaselike.Execute(ctx, StoryId)
				if err != nil {
					log.Println(err, "failed to increase like count")
				}
			case common.ReactedStoryUnlike:
				_, err := sbr.decreaselike.Execute(ctx, StoryId)
				if err != nil {
					log.Println(err, "failed to decrease like count")
				}

			default:
				log.Println("unknown react type")

			}

		}
	}()
}
