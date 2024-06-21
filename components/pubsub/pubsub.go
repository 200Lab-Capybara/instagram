package pubsub

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MessageBroker interface {
	Publish(ctx context.Context, topic string, msg *AppMessage) error
	Subscribe(ctx context.Context, topic string) (subChan <-chan *AppMessage, close func(), err error)
}

type AppMessage struct {
	Id        uuid.UUID
	SenderId  *uuid.UUID
	Topic     string
	Payload   map[string]interface{}
	CreatedAt time.Time
}

func NewAppMessage(senderId *uuid.UUID, topic string, payload map[string]interface{}) *AppMessage {
	id, _ := uuid.NewV7()
	return &AppMessage{Id: id, SenderId: senderId, Topic: topic, Payload: payload, CreatedAt: time.Now()}
}

func NewSimpleAppMessage(payload map[string]interface{}) *AppMessage {
	return NewAppMessage(nil, "", payload)
}
