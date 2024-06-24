package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	reactstorymysql "instagram/app/infras/services/reaction_story/repository/mysql"
	"instagram/app/infras/services/reaction_story/repository/rpc_client"
	reactionstoryhttp "instagram/app/infras/services/reaction_story/transport/http"
	"instagram/app/internals/services/reaction_story/usecase"
	"instagram/common"
	"instagram/components/pubsub/natspubsub"
)

func BuildReactStoryService(con common.SQLDatabase, nat *nats.Conn, v1 *gin.RouterGroup) {

	reactionRepo := reactstorymysql.NewMySQLStorage(con)
	getStoryRepo := rpc_client.NewGetStoryRepo(con)
	ps := natspubsub.NewNatsProvider(nat)
	getReactionStory := reactionstoryusecase.NewInsertReactionStoryUseCase(reactionRepo, getStoryRepo, ps)
	reactionStoryHandler := reactionstoryhttp.NewReactionStoryHandler(getReactionStory)
	reactionStoryHandler.RegisterV1Router(v1)
}
