package followhttp

import (
	"github.com/gin-gonic/gin"
	followuserusecase "instagram/app/internals/services/follow/usecase"
)

type followUserHandler struct {
	followUserUseCase   followuserusecase.FollowUserUseCase
	getFollowingUseCase followuserusecase.GetListFollowingUseCase
	getFollowerUseCase  followuserusecase.GetListFollowerUseCase
}

func NewFollowUserHandler(followUserUseCase followuserusecase.FollowUserUseCase, getFollowingUseCase followuserusecase.GetListFollowingUseCase, getFollowerUseCase followuserusecase.GetListFollowerUseCase) *followUserHandler {
	return &followUserHandler{
		followUserUseCase:   followUserUseCase,
		getFollowingUseCase: getFollowingUseCase,
		getFollowerUseCase:  getFollowerUseCase,
	}
}

func (hdl *followUserHandler) RegisterV1Router(v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	v1.POST("/follow/:following_id", middleware, hdl.FollowUnfollowUser())
	v1.GET("/follow/:user_id/following", middleware, hdl.GetFollowingHandler())
	v1.GET("/follow/:user_id/follower", middleware, hdl.GetListFollowerHandler())
}
