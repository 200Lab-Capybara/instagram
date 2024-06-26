package reactionstoryhttp

import (
	"github.com/gin-gonic/gin"
	"instagram/app/internals/services/reaction_story/usecase"
)

type reactionStoryHandler struct {
	uc reactionstoryusecase.InsertReactionStoryUseCase
}

func NewReactionStoryHandler(uc reactionstoryusecase.InsertReactionStoryUseCase) *reactionStoryHandler {
	return &reactionStoryHandler{uc: uc}
}

func (hdl *reactionStoryHandler) RegisterV1Router(r *gin.RouterGroup, middleware gin.HandlerFunc) {
	r.POST("/story/:id/reaction", middleware, hdl.ReactStoryHandler())
}
