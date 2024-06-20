package reactionstoryusecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_story/model"
)

type reactionStoryUC struct {
	reactionStoryRepo IReactionStoryRepository
	storyRepo         getStoryRepository
}

func NewInsertReactionStoryUserCase(reactRepo IReactionStoryRepository, storyRepo getStoryRepository) InsertReactionStoryUserCase {
	return &reactionStoryUC{
		reactRepo,
		storyRepo,
	}
}

func (u *reactionStoryUC) Execute(ctx context.Context, storyId uuid.UUID, userId uuid.UUID) (bool, error) {
	existStory, err := u.storyRepo.FindStoryById(ctx, storyId)
	if err != nil {
		if !errors.Is(err, model.StoryNotFound) {
			return false, model.StoryNotFound
		}
		return false, err
	}
	if existStory == nil {
		return false, model.StoryNotFound
	}

	existReactStory, err := u.reactionStoryRepo.HasBeenReactionStory(ctx, storyId, userId)
	if err != nil {
		return false, err
	}
	if existReactStory != nil {
		_, err = u.reactionStoryRepo.DelReactionStory(ctx, storyId, userId)
		if err != nil {
			return false, err
		}
	} else {
		_, err = u.reactionStoryRepo.CreateNewReactionStory(ctx, storyId, userId)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

type InsertReactionStoryUserCase interface {
	Execute(ctx context.Context, storyId uuid.UUID, userId uuid.UUID) (bool, error)
}

type IReactionStoryRepository interface {
	CreateNewReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
	HasBeenReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (*model.ReactionStory, error)
	DelReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
}
type getStoryRepository interface {
	FindStoryById(ctx context.Context, sid uuid.UUID) (*model.Story, error)
}
