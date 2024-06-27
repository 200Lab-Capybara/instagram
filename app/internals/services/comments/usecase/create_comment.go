package usecasecomment

import (
	"context"
	"github.com/google/uuid"
	modelcomment "instagram/app/internals/services/comments/model"
	"instagram/common"
	"time"
)

//1. check post xem co ton tai hay ko
//2. create comment trong post
//3. public qua post

type CreateCommentUseCae interface {
	Execute(ctx context.Context, requester common.Requester, dto modelcomment.CommentCreation) (*uuid.UUID, error)
}

type CreateCommentRepository interface {
	CreateComment(ctx context.Context, dto modelcomment.Comment) (*uuid.UUID, error)
}

type CheckPostRepository interface {
	CheckExistPostById(ctx context.Context, id uuid.UUID) (bool, error)
}

type commentUseCae struct {
	createCommentRepo CreateCommentRepository
	checkPostRepo     CheckPostRepository
}

//1. check post xem co ton tai hay ko
//1.1 ko ton tai -> true
//1.2 co ton tai -> false ()
//2. create comment trong post
//3. public qua post

func (c *commentUseCae) Execute(ctx context.Context, requester common.Requester, dto modelcomment.CommentCreation) (*uuid.UUID, error) {
	//1. check post
	exist, err := c.checkPostRepo.CheckExistPostById(ctx, dto.PostId)
	if err != nil {
		return nil, err
	}
	//1.1 ko ton tai
	if !exist {
		return nil, common.NewCustomError(modelcomment.ErrPostNotFound, modelcomment.ErrPostNotFound.Error(), "post_not_found")
	}
	//1.2 co ton tai
	//2. create comment
	commentId, _ := uuid.NewV7()

	commentDTO := modelcomment.Comment{
		Id:         commentId,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Content:    dto.Content,
		UserId:     requester.UserId(),
		PostId:     dto.PostId,
		ReactCount: 0,
	}

	_, err = c.createCommentRepo.CreateComment(ctx, commentDTO)
	if err != nil {
		return nil, err
	}
	return &commentId, nil
}

func NewCommentUseCase(createCommentRepo CreateCommentRepository, checkPostRepo CheckPostRepository) CreateCommentUseCae {
	return &commentUseCae{
		createCommentRepo: createCommentRepo,
		checkPostRepo:     checkPostRepo,
	}
}
