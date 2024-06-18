package reactionposthttp

import (
	"github.com/gin-gonic/gin"
	reactionpostusecase "github.com/nghiatrann0502/instagram-clone/app/internals/services/reaction_post/usecase"
)

type reactionPostHandler struct {
	uc reactionpostusecase.LikePostUseCase
}

func NewReactionPostHandler(uc reactionpostusecase.LikePostUseCase) *reactionPostHandler {
	return &reactionPostHandler{uc: uc}
}

func (hdl *reactionPostHandler) RegisterV1Router(r *gin.RouterGroup) {
	r.POST("/posts/:id/like", hdl.ReactPostHandler())
}
