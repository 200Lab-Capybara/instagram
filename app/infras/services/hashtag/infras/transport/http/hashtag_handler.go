package hashtaghttp

import (
	"github.com/gin-gonic/gin"
	hashtagusercase "instagram/app/internals/services/hashtag/usecase"
)

type hashtagHandler struct {
	uc hashtagusercase.AddingHashTagUseCase
}

//	func NewReactionPostHandler(uc reactionpostusecase.LikePostUseCase) *reactionPostHandler {
//		return &reactionPostHandler{uc: uc}
//	}
func NewHashTagHandler(uc hashtagusercase.AddingHashTagUseCase) *hashtagHandler {
	return &hashtagHandler{uc: uc}
}

//func (hdl *reactionPostHandler) RegisterV1Router(r *gin.RouterGroup) {
//	r.POST("/posts/:id/like", hdl.ReactPostHandler())
//}

func (hdl *hashtagHandler) RegisterV1Router(r *gin.RouterGroup) {
	r.POST("/posts/:id/hashtag", hdl.HashTagHandler())
}
