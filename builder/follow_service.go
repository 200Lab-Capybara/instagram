package builder

import (
	"github.com/gin-gonic/gin"
	followmysql "instagram/app/infras/services/follow/repository/mysql"
	followrpc_client "instagram/app/infras/services/follow/repository/rpc_client"
	followhttp "instagram/app/infras/services/follow/transport/http"
	followusecase "instagram/app/internals/services/follow/usecase"
	"instagram/components/pubsub/natspubsub"
)

func BuildFollowService(ctxSvr ServiceContext, middleware gin.HandlerFunc) {
	followStore := followmysql.NewMysqlStorage(ctxSvr.GetDB())
	pubsub := natspubsub.NewNatsProvider(ctxSvr.GetNatsConn())

	rpcConClient := BuildUserRpcClient()
	rpcClient := followrpc_client.NewFollowRpcClient(rpcConClient)

	followUnfollowUC := followusecase.NewFollowUserUseCase(followStore, pubsub)
	getFollowingUC := followusecase.NewGetListFollowingUseCase(followStore, rpcClient)
	getFollowerUC := followusecase.NewGetListFollowerUseCase(rpcClient, followStore)
	getInternalFollowingUC := followusecase.NewGetInternalFollowingUseCase(followStore)

	followHandler := followhttp.NewFollowUserHandler(followUnfollowUC, getFollowingUC, getFollowerUC, getInternalFollowingUC)
	followHandler.RegisterV1Router(ctxSvr.GetV1(), middleware)
	followHandler.RegisterInternalRouter(ctxSvr.GetV1Internal(), middleware)
}
