package reactionposthttp

import (
	"github.com/gin-gonic/gin"
	reactionpostusecase "instagram/app/internals/services/reaction_post/usecase"
)

type reactionPostHandler struct {
	reactionPostUC    reactionpostusecase.ReactionPostUseCase
	getUserLikePostUC reactionpostusecase.GetUserLikePostUseCase
}

func NewReactionPostHandler(reactionPostUC reactionpostusecase.ReactionPostUseCase, getUserLikePostUC reactionpostusecase.GetUserLikePostUseCase) *reactionPostHandler {
	return &reactionPostHandler{
		reactionPostUC:    reactionPostUC,
		getUserLikePostUC: getUserLikePostUC,
	}
}

func (hdl *reactionPostHandler) RegisterV1Router(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	r.POST("/posts/:id/react", middleware, hdl.ReactPostHandler())
	r.GET("/posts/:id/likes", middleware, hdl.GetUserLikePostHandler())
}
