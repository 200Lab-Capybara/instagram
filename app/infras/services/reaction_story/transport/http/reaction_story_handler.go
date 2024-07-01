package reactionstoryhttp

import (
	"github.com/gin-gonic/gin"
	"instagram/app/internals/services/reaction_story/usecase"
)

type reactionStoryHandler struct {
	reactionStoryUseCase reactionstoryusecase.ReactionStoryUseCase
}

func NewReactionStoryHandler(uc reactionstoryusecase.ReactionStoryUseCase) *reactionStoryHandler {
	return &reactionStoryHandler{reactionStoryUseCase: uc}
}

func (hdl *reactionStoryHandler) RegisterV1Router(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	r.POST("/story/:id/reaction", middleware, hdl.ReactStoryHandler())
}
