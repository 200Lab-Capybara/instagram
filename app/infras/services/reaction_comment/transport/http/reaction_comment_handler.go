package httpreactioncomment

import (
	"github.com/gin-gonic/gin"
	usecasereactioncomment "instagram/app/internals/services/reaction_comment/usecase"
)

type reactionCommentHandler struct {
	uc usecasereactioncomment.InsertReactionCommentUseCase
}

func NewReactionCommentHandler(uc usecasereactioncomment.InsertReactionCommentUseCase) *reactionCommentHandler {
	return &reactionCommentHandler{uc: uc}
}
func (hdl *reactionCommentHandler) RegisterV1Router(r *gin.RouterGroup) {
	r.POST("/comment/:id/reaction", hdl.ReactCommentHandler())
}
