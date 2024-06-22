package pubsub

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type MessageBroker interface {
	Publish(ctx context.Context, msg *Message) error
	Subscribe(ctx context.Context, topic string) (subChan <-chan *Message, close func(), err error)
}

//type AppMessage interface {
//	GetID() uuid.UUID
//	GetSenderId() *uuid.UUID
//	GetTopic() string
//	GetPayload() map[string]interface{}
//	GetCreatedAt() time.Time
//}

type Message struct {
	ID        uuid.UUID
	SenderID  *uuid.UUID
	Topic     string
	Payload   map[string]any
	CreatedAt time.Time
}

func NewAppMessage(senderID *uuid.UUID, topic string, payload map[string]any) *Message {
	id, _ := uuid.NewV7()
	return &Message{
		ID:        id,
		SenderID:  senderID,
		Topic:     topic,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}

func NewSimpleAppMessage(payload map[string]any) *Message {
	return NewAppMessage(nil, "", payload)
}

//
//func (a *Message) GetID() uuid.UUID {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Message) GetSenderId() *uuid.UUID {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Message) GetTopic() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Message) GetPayload() map[string]interface{} {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (a *Message) GetCreatedAt() time.Time {
//	//TODO implement me
//	panic("implement me")
//}
