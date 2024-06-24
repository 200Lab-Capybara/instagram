package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	reactionpostmysql "instagram/app/infras/services/reaction_post/repository/mysql"
	"instagram/app/infras/services/reaction_post/repository/rpc_client"
	reactionposthttp "instagram/app/infras/services/reaction_post/transport/http"
	reactionpostusecase "instagram/app/internals/services/reaction_post/usecase"
	"instagram/common"
	"instagram/components/pubsub/natspubsub"
)

func BuildReactPostService(con common.SQLDatabase, nat *nats.Conn, v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	reactionRepo := reactionpostmysql.NewMySQLStorage(con)
	getPostRepo := rpc_client.NewGetPostRepo(con)
	ps := natspubsub.NewNatsProvider(nat)
	reactionUC := reactionpostusecase.NewLikePostUseCase(reactionRepo, getPostRepo, ps)
	reactionHDL := reactionposthttp.NewReactionPostHandler(reactionUC)
	reactionHDL.RegisterV1Router(v1, middleware)
}
