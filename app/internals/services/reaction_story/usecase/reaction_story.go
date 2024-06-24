package reactionstoryusecase

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_story/model"
)

type reactionStoryUC struct {
	reactionStoryRepo IReactionStoryRepository
	storyRepo         getStoryRepository
}

func NewInsertReactionStoryUseCase(reactRepo IReactionStoryRepository, storyRepo getStoryRepository) InsertReactionStoryUseCase {
	return &reactionStoryUC{
		reactRepo,
		storyRepo,
	}
}

func (u *reactionStoryUC) Execute(ctx context.Context, storyId uuid.UUID, userId uuid.UUID) (bool, error) {
	_, err := u.storyRepo.FindStoryById(ctx, storyId)
	if err != nil {
		return false, err
	}

	existReactStory, err := u.reactionStoryRepo.HasBeenReactionStory(ctx, storyId, userId)
	if existReactStory && err == nil {
		_, err = u.reactionStoryRepo.RemoveReactionStory(ctx, storyId, userId)
		if err != nil {
			return false, err
		}
		_, err = u.storyRepo.DecreaseReactCountById(ctx, storyId)
		if err != nil {
			return false, err
		}
	} else {
		_, err = u.reactionStoryRepo.CreateNewReactionStory(ctx, storyId, userId)
		if err != nil {
			return false, err
		}
		_, err = u.storyRepo.IncreaseReactCountById(ctx, storyId)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

type InsertReactionStoryUseCase interface {
	Execute(ctx context.Context, storyId uuid.UUID, userId uuid.UUID) (bool, error)
}

type IReactionStoryRepository interface {
	CreateNewReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
	HasBeenReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
	RemoveReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error)
}
type getStoryRepository interface {
	FindStoryById(ctx context.Context, sid uuid.UUID) (*model.Story, error)
	IncreaseReactCountById(ctx context.Context, sid uuid.UUID) (bool, error)
	DecreaseReactCountById(ctx context.Context, sid uuid.UUID) (bool, error)
}
