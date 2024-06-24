package builder

import (
	"github.com/gin-gonic/gin"
	nats2 "github.com/nats-io/nats.go"
	postsmysql "instagram/app/infras/services/posts/repository/mysql"
	postrpc_client "instagram/app/infras/services/posts/repository/rpc_client"
	postshttp "instagram/app/infras/services/posts/transport/http"
	postsubscriber "instagram/app/infras/services/posts/transport/subscriber"
	postusecase "instagram/app/internals/services/posts/usecase"
	"instagram/common"
	"instagram/components/pubsub/natspubsub"
)

func BuildPostService(con common.SQLDatabase, v1 *gin.RouterGroup, pubsubCon *nats2.Conn, middleware gin.HandlerFunc) {
	postStorage := postsmysql.NewMysqlStorage(con)
	pubsub := natspubsub.NewNatsProvider(pubsubCon)

	rpcConClient := BuildUserRpcClient()
	rpcClient := postrpc_client.NewPostRpcClient(rpcConClient)

	createPostUseCase := postusecase.NewCreatePostUseCase(postStorage, pubsub)
	getListPostsByUserIdUC := postusecase.NewGetListPostByUserIdUseCase(postStorage, rpcClient)
	increaseLikeCountUC := postusecase.NewIncreaseLikeCountUseCase(postStorage)
	decreaseLikeCountUC := postusecase.NewDecreaseLikeCountUseCase(postStorage)

	// RPC Client

	// HTTP Handler
	postHandler := postshttp.NewPostHandler(createPostUseCase, getListPostsByUserIdUC)
	postHandler.RegisterV1Router(v1, middleware)

	// Post Subscriber
	postSubscriber := postsubscriber.NewPostSubscriber(pubsub, increaseLikeCountUC, decreaseLikeCountUC)
	postSubscriber.Init()
}
