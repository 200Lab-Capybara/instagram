package reactionstoryusecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_story/model"
	"instagram/common"
	"instagram/components/pubsub"
	"time"
)

type reactionStoryUC struct {
	reactionStoryRepo IReactionStoryRepository
	storyRepo         getStoryRepository
	pubsub            pubsub.MessageBroker
}

func NewInsertReactionStoryUseCase(reactRepo IReactionStoryRepository, storyRepo getStoryRepository, pubsub pubsub.MessageBroker) ReactionStoryUseCase {
	return &reactionStoryUC{
		reactRepo,
		storyRepo,
		pubsub,
	}
}

func (u *reactionStoryUC) Execute(ctx context.Context, storyId uuid.UUID, request common.Requester) (bool, error) {
	userId := request.UserId()
	_, err := u.storyRepo.FindStoryById(ctx, storyId)
	if err != nil {
		return false, err
	}
	reactType := common.ReactedStoryLike
	storyModel := model.ReactionStory{
		UserId:    userId,
		StoryId:   storyId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	existReactStory, err := u.reactionStoryRepo.HasBeenReactionStory(ctx, storyModel)
	if existReactStory && err == nil {
		_, err = u.reactionStoryRepo.RemoveReactionStory(ctx, storyId, userId)
		if err != nil {
			return false, err
		}
		reactType = common.ReactedStoryUnlike
	} else {
		_, err = u.reactionStoryRepo.CreateNewReactionStory(ctx, storyId, userId)
		if err != nil {
			return false, err
		}
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Error public message from topic %s", common.ReactedStoryTopic)
			}
		}()

		postMessage := pubsub.NewAppMessage(&userId, common.ReactedStoryTopic, map[string]interface{}{
			"story_id":   storyId,
			"react_type": reactType,
		})

		err := u.pubsub.Publish(ctx, postMessage)
		if err != nil {
			panic(err)
		}
	}()
	return true, nil
}

type ReactionStoryUseCase interface {
	Execute(ctx context.Context, storyId uuid.UUID, request common.Requester) (bool, error)
}

type IReactionStoryRepository interface {
	CreateNewReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
	HasBeenReactionStory(ctx context.Context, story model.ReactionStory) (bool, error)
	RemoveReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
}
type getStoryRepository interface {
	FindStoryById(ctx context.Context, sid uuid.UUID) (*model.Story, error)
	//IncreaseReactCountById(ctx context.Context, sid uuid.UUID) (bool, error)
	//DecreaseReactCountById(ctx context.Context, sid uuid.UUID) (bool, error)
}
