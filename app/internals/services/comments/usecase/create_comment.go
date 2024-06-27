package usecasecomment

import (
	"context"
	modelcomment "instagram/app/internals/services/comments/model"
	"instagram/common"
)

//1. check post xem co ton tai hay ko
//2. create comment trong post
//3. public qua post

type CreateCommentUseCae interface {
	Execute(ctx context.Context, requester common.Requester, dto modelcomment.CommentCreation) (string, error)
}

type commentUseCae struct {
}

func (c *commentUseCae) Execute(ctx context.Context, requester common.Requester, dto modelcomment.CommentCreation) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewCommentUseCase() CreateCommentUseCae {
	return &commentUseCae{}
}
