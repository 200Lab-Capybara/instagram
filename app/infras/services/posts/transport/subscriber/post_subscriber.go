package postsubscriber

import (
	"context"
	"fmt"
	"instagram/common"
	"instagram/components/pubsub"
)

type postSubscriber struct {
	pubsub pubsub.MessageBroker
}

func NewPostSubscriber(pubsub pubsub.MessageBroker) *postSubscriber {
	return &postSubscriber{pubsub: pubsub}
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
			fmt.Println(msg.Payload)
		}
	}()
}
