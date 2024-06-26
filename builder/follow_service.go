package builder

import (
	"github.com/gin-gonic/gin"
	followmysql "instagram/app/infras/services/follow/repository/mysql"
	followhttp "instagram/app/infras/services/follow/transport/http"
	followusecase "instagram/app/internals/services/follow/usecase"
	"instagram/components/pubsub/natspubsub"
)

func BuildFollowService(ctxSvr ServiceContext, middleware gin.HandlerFunc) {
	followStore := followmysql.NewMysqlStorage(ctxSvr.GetDB())
	pubsub := natspubsub.NewNatsProvider(ctxSvr.GetNatsConn())

	followUnfollowUC := followusecase.NewFollowUserUseCase(followStore, pubsub)

	followHandler := followhttp.NewFollowUserHandler(followUnfollowUC)
	followHandler.RegisterV1Router(ctxSvr.GetV1(), middleware)
}
