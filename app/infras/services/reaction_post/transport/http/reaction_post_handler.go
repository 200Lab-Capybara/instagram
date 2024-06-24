package reactionposthttp

import (
	"github.com/gin-gonic/gin"
	reactionpostusecase "instagram/app/internals/services/reaction_post/usecase"
)

type reactionPostHandler struct {
	uc reactionpostusecase.ReactionPostUseCase
}

func NewReactionPostHandler(uc reactionpostusecase.ReactionPostUseCase) *reactionPostHandler {
	return &reactionPostHandler{uc: uc}
}

func (hdl *reactionPostHandler) RegisterV1Router(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	r.POST("/posts/:id/react", middleware, hdl.ReactPostHandler())
}
