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

type GetCommentRepository interface {
	FindCommentById(ctx context.Context, commentId uuid.UUID) (*modelreactioncomment.Comment, error)
	IncreaseReactionCountById(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
	DecreaseReactionCountById(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
}

type reactionCommentUC struct {
	reactionCommentRepo IReactionCommentRepository
	commentRepo         GetCommentRepository
}

type InsertReactionCommentUseCase interface {
	Execute(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error)
}

func NewInsertReactionCommentUseCase(reactRepo IReactionCommentRepository, commentRepo GetCommentRepository) InsertReactionCommentUseCase {
	return &reactionCommentUC{
		reactRepo,
		commentRepo,
	}
}

func (u *reactionCommentUC) Execute(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error) {
	// 1. Tim comment voi commentId xem co ko
	// 2. Check comment da co react chua
	// 2.1 Neu da co react thi remove react
	// 2.1.1 Sau do Decrease react
	// 2.2 Neu chua co thi create react
	// 2.2.2 Sau do Increase react

	// 1
	_, err := u.commentRepo.FindCommentById(ctx, commentId)
	if err != nil {
		return false, err
	}

	// 2
	existed, err := u.reactionCommentRepo.HasBeenReactionComment(ctx, commentId, userId)
	if existed == true {
		// 2.1
		_, err := u.reactionCommentRepo.RemoveReactionComment(ctx, commentId, userId)
		if err != nil {
			return false, err
		}
		// 2.1.1
		_, err = u.commentRepo.DecreaseReactionCountById(ctx, commentId, userId)
		if err != nil {
			return false, err
		}

	} else {
		//2.2
		_, err := u.reactionCommentRepo.CreateNewReactionComment(ctx, commentId, userId)
		if err != nil {
			return false, err
		}
		//2.2.2
		_, err = u.commentRepo.IncreaseReactionCountById(ctx, commentId, userId)
		if err != nil {
			return false, err
		}
	}

	return true, nil

}
