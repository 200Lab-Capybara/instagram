package followhttp

import (
	"github.com/gin-gonic/gin"
	followuserusecase "instagram/app/internals/services/follow/usecase"
)

type followUserHandler struct {
	followUserUseCase followuserusecase.FollowUserUseCase
}

func NewFollowUserHandler(followUserUseCase followuserusecase.FollowUserUseCase) *followUserHandler {
	return &followUserHandler{
		followUserUseCase: followUserUseCase,
	}
}

func (hdl *followUserHandler) RegisterV1Router(v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	v1.POST("/follow/:following_id", middleware, hdl.FollowUnfollowUser())
}
