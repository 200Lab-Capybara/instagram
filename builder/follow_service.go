package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	followmysql "instagram/app/infras/services/follow/repository/mysql"
	followhttp "instagram/app/infras/services/follow/transport/http"
	followusecase "instagram/app/internals/services/follow/usecase"
	"instagram/common"
	"instagram/components/pubsub/natspubsub"
)

func BuildFollowService(con common.SQLDatabase, v1 *gin.RouterGroup, pubsubCon *nats.Conn, middleware gin.HandlerFunc) {
	followStore := followmysql.NewMysqlStorage(con)
	pubsub := natspubsub.NewNatsProvider(pubsubCon)

	followUnfollowUC := followusecase.NewFollowUserUseCase(followStore, pubsub)

	followHandler := followhttp.NewFollowUserHandler(followUnfollowUC)
	followHandler.RegisterV1Router(v1, middleware)
}
