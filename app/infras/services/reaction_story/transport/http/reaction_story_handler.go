package reactionstoryhttp

import (
	"github.com/gin-gonic/gin"
	"instagram/app/internals/services/reaction_story/usecase"
)

type reactionStoryHandler struct {
	uc reactionstoryusecase.InsertReactionStoryUserCase
}

func NewReactionStoryHandler(uc reactionstoryusecase.InsertReactionStoryUserCase) *reactionStoryHandler {
	return &reactionStoryHandler{uc: uc}
}

func (hdl *reactionStoryHandler) RegisterV1Router(r *gin.RouterGroup) {
	r.POST("/story/:id/reaction", hdl.ReactStoryHandler())
}
