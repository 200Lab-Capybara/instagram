package storyhttp

import (
	"github.com/gin-gonic/gin"
	"instagram/app/internals/services/stories/usecase"
)

type storyHandler struct {
	storyUC storyusecase.CreateStoryUC
}

func NewPostHandler(storyUC storyusecase.CreateStoryUC) *storyHandler {
	return &storyHandler{
		storyUC: storyUC,
	}
}

func (hdl *storyHandler) RegisterV1Router(v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	v1.POST("/story", middleware, hdl.CreateStoryHandler())
}
