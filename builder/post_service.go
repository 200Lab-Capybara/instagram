package builder

import (
	"github.com/gin-gonic/gin"
	postapi "instagram/app/infras/services/posts/repository/api"
	postsmysql "instagram/app/infras/services/posts/repository/mysql"
	postrpc_client "instagram/app/infras/services/posts/repository/rpc_client"
	postshttp "instagram/app/infras/services/posts/transport/http"
	postsubscriber "instagram/app/infras/services/posts/transport/subscriber"
	postusecase "instagram/app/internals/services/posts/usecase"
	"instagram/components/pubsub/natspubsub"
)

func BuildPostService(svr ServiceContext, middleware gin.HandlerFunc) {
	con := svr.GetDB()
	pubsubCon := svr.GetNatsConn()
	v1 := svr.GetV1()

	postStorage := postsmysql.NewMysqlStorage(con)
	pubsub := natspubsub.NewNatsProvider(pubsubCon)

	rpcConClient := BuildUserRpcClient()
	rpcClient := postrpc_client.NewPostRpcClient(rpcConClient)

	followAPI := postapi.NewFollowAPI("http://localhost:8001/internal/v1/rpc")

	createPostUseCase := postusecase.NewCreatePostUseCase(postStorage, pubsub)
	getListPostsByUserIdUC := postusecase.NewGetListPostByUserIdUseCase(postStorage, rpcClient)
	increaseLikeCountUC := postusecase.NewIncreaseLikeCountUseCase(postStorage)
	decreaseLikeCountUC := postusecase.NewDecreaseLikeCountUseCase(postStorage)
	getFeedUC := postusecase.NewGetFeedUseCase(rpcClient, postStorage, followAPI)

	// RPC Client

	// HTTP Handler
	postHandler := postshttp.NewPostHandler(createPostUseCase, getListPostsByUserIdUC, getFeedUC)
	postHandler.RegisterV1Router(v1, middleware)

	// Post Subscriber
	postSubscriber := postsubscriber.NewPostSubscriber(pubsub, increaseLikeCountUC, decreaseLikeCountUC)
	postSubscriber.Init()
}
