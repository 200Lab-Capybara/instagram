package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"instagram/app/internals/services/stories/model"
	"instagram/common"
	"instagram/components/pubsub"
	"time"
)

type createStoryUC struct {
	createstoryuc CreateStoryRepo
	pubsub        pubsub.MessageBroker
}

func NewCreateStoryUC(createstory CreateStoryRepo, pubsub pubsub.MessageBroker) CreateStoryUC {
	return &createStoryUC{
		createstory,
		pubsub,
	}
}

func (uc createStoryUC) Execute(ctx context.Context, requester common.Requester, dto model.CreateStory) (*uuid.UUID, error) {
	userId := requester.UserId()
	storyId, _ := uuid.NewV7()
	imageId, _ := uuid.NewV7()

	story := &model.Story{
		Id:      storyId,
		UserId:  userId,
		Content: dto.Content,
		Image: model.ImageStory{
			Id:          imageId,
			UserId:      userId,
			ImageUrl:    dto.ImageUrl,
			Size:        dto.Size,
			Width:       dto.Width,
			Height:      dto.Height,
			Status:      common.ImageUsed,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			StorageName: dto.Storage,
		},
		Count:       0,
		IsActive:    true,
		ExpiresTime: 24,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	id, err := uc.createstoryuc.CreateStory(ctx, story)
	if err != nil {
		return nil, err
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Error public message from topic %s", common.CreatedStoryTopic)
			}
		}()

		postMessage := pubsub.NewAppMessage(&userId, common.CreatedStoryTopic, map[string]interface{}{
			"story_id": story,
			"user_id":  userId,
		})

		// TODO: Publish CreatedPostTopic event
		err := uc.pubsub.Publish(ctx, postMessage)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Error public message from topic %s", common.CreatedImageTopic)
			}
		}()

		postMessage := pubsub.NewAppMessage(&userId, common.CreatedImageTopic, map[string]interface{}{
			"image":       story.Image,
			"create_type": common.CreateStoryImage,
		})

		// TODO: Publish CreatedPostTopic event
		err := uc.pubsub.Publish(ctx, postMessage)
		if err != nil {
			panic(err)
		}
	}()

	return id, nil
}

type CreateStoryUC interface {
	Execute(ctx context.Context, requester common.Requester, dto model.CreateStory) (*uuid.UUID, error)
}
type CreateStoryRepo interface {
	CreateStory(ctx context.Context, story *model.Story) (*uuid.UUID, error)
}
