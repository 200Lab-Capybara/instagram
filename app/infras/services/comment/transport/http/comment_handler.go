package httpcomment

import (
	"github.com/gin-gonic/gin"
	usecasecomment "instagram/app/internals/services/comments/usecase"
)

type commentHandler struct {
	createCommentUC usecasecomment.CreateCommentUseCae
}

func NewCommentHandler(createCommentUC usecasecomment.CreateCommentUseCae) *commentHandler {
	return &commentHandler{
		createCommentUC: createCommentUC,
	}
}

func (hdl *commentHandler) CommentV1Router(v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	v1.POST("/posts/:id/comment", middleware, hdl.CreateCommentHandler())
}
