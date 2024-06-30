package postsubscriber

import (
	"context"
	"github.com/google/uuid"
	postusecase "instagram/app/internals/services/posts/usecase"
	"instagram/common"
	"instagram/components/pubsub"
	"log"
)

type postSubscriber struct {
	pubsub              pubsub.MessageBroker
	increaseLikeCount   postusecase.IncreaseLikeCountUseCase
	decreaseLikeCountUC postusecase.DecreaseLikeCountUseCase
}

func NewPostSubscriber(pubsub pubsub.MessageBroker, increaseLikeCount postusecase.IncreaseLikeCountUseCase, decreaseLikeCountUC postusecase.DecreaseLikeCountUseCase) *postSubscriber {
	return &postSubscriber{
		pubsub:              pubsub,
		increaseLikeCount:   increaseLikeCount,
		decreaseLikeCountUC: decreaseLikeCountUC,
	}
}

func (sbr *postSubscriber) Init() {
	go func() {
		ctx := context.Background()
		subscribe, c, err := sbr.pubsub.Subscribe(ctx, common.ReactedPostTopic)
		if err != nil {
			return
		}

		defer c()

		for msg := range subscribe {
			if msg == nil {
				return
			}

			payload := msg.Payload
			postID := uuid.MustParse(payload["post_id"].(string))
			reactType := payload["react_type"]

			switch reactType {
			case common.ReactedPostLike:
				_, err := sbr.increaseLikeCount.Execute(ctx, postID)
				if err != nil {
					log.Println(err, "failed to increase like count")
				}
			case common.ReactedPostUnlike:
				_, err := sbr.decreaseLikeCountUC.Execute(ctx, postID)
				if err != nil {
					log.Println(err, "failed to decrease like count")
				}

			default:
				log.Println("unknown react type")

			}

		}
	}()

}
