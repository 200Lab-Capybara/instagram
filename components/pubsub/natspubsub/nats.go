package natspubsub

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"instagram/components/pubsub"
	"log"
)

type natsProvider struct {
	con *nats.Conn
}

func NewNatsProvider(con *nats.Conn) pubsub.MessageBroker {
	return &natsProvider{
		con: con,
	}
}

func (n *natsProvider) Publish(ctx context.Context, msg *pubsub.Message) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = n.con.Publish(msg.Topic, bytes)

	if err != nil {
		return err
	}

	return nil
}

func (n *natsProvider) Subscribe(ctx context.Context, topic string) (subChan <-chan *pubsub.Message, close func(), err error) {
	ch := make(chan *pubsub.Message)
	sub, err := n.con.Subscribe(topic, func(msg *nats.Msg) {
		var data pubsub.Message
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Println(err)
		}

		message := pubsub.NewAppMessage(data.SenderID, topic, data.Payload)

		ch <- message
	})

	close = func() {
		err := sub.Unsubscribe()
		if err != nil {
			log.Fatal(err)
		}
	}

	return ch, close, nil
}
