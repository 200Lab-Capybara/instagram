package usecasereactioncomment

import (
	"context"
	"github.com/google/uuid"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
)

type IReactionCommentRepository interface {
	CreateNewReactionComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
	HasBeenReactionComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
	RemoveReactionComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
}

type getCommentRepository interface {
	FinCommentById(ctx context.Context, commentId uuid.UUID) (modelreactioncomment.Comment, error)
	IncreaseReactionCountById(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
	DecreaseReactionCountById(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
}

type reactionCommentUC struct {
	reactionCommentRepo IReactionCommentRepository
	commentRepo         getCommentRepository
}

type InsertReactionCommentUseCase interface {
	Execute(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
}

func NewInsertReactionCommentUseCase(reactRepo IReactionCommentRepository, commentRepo getCommentRepository) InsertReactionCommentUseCase {
	return &reactionCommentUC{
		reactRepo,
		commentRepo,
	}
}

func (u *reactionCommentUC) Execute(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error) {
	_, err := u.reactionCommentRepo.CreateNewReactionComment(ctx, commentId, userId)
	if err != nil {
		return false, err
	}

	existReactionComment, err := u.reactionCommentRepo.HasBeenReactionComment(ctx, commentId, userId)
	if err != nil {
		return false, err
	}

	if existReactionComment {
		_, err = u.reactionCommentRepo.RemoveReactionComment(ctx, commentId, userId)
		if err != nil {
			return false, err
		}
		_, err = u.reactionCommentRepo.CreateNewReactionComment(ctx, commentId, userId)
	} else {
		_, err = u.commentRepo.IncreaseReactionCountById(ctx, commentId, userId)
		if err != nil {
			return false, err
		}
		_, err = u.commentRepo.DecreaseReactionCountById(ctx, commentId, userId)
	}

	return true, nil

}
