package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_story/model"
)

type reactionStoryUC struct {
	reactionStoryRepo reactionStoryRepository
	storyRepo         getStoryRepository
}

func NewInsertReactionStoryUserCase(reactRepo reactionStoryRepository, storyRepo getStoryRepository) InsertReactionStoryUserCase {
	return &reactionStoryUC{
		reactRepo,
		storyRepo,
	}
}

func (u *reactionStoryUC) Execute(ctx context.Context, story *model.ReactionStory) (bool, error) {
	existStory, err := u.storyRepo.FindStoryById(ctx, story.StoryId)
	if err != nil {
		if !errors.Is(err, model.StoryNotFound) {
			return false, model.StoryNotFound
		}
		return false, err
	}
	if existStory == nil {
		return false, model.StoryNotFound
	}

	existReactStory, err := u.reactionStoryRepo.hasBeenReactionStory(ctx, story.StoryId, story.UserId)
	if err != nil {
		return false, err
	}
	if existReactStory != nil {
		_, err = u.reactionStoryRepo.delReactionStory(ctx, story.StoryId, story.UserId)
		if err != nil {
			return false, err
		}
	} else {
		_, err = u.reactionStoryRepo.createNewReactionStory(ctx, story.StoryId, story.UserId)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

type InsertReactionStoryUserCase interface {
	Execute(ctx context.Context, user *model.ReactionStory) (bool, error)
}

type reactionStoryRepository interface {
	createNewReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
	hasBeenReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (*model.ReactionStory, error)
	delReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
}
type getStoryRepository interface {
	FindStoryById(ctx context.Context, sid uuid.UUID) (*model.Story, error)
}
