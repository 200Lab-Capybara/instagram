package builder

import (
	"github.com/gin-gonic/gin"
	reactstorymysql "instagram/app/infras/services/reaction_story/repository/mysql"
	"instagram/app/infras/services/reaction_story/repository/rpc_client"
	reactionstoryhttp "instagram/app/infras/services/reaction_story/transport/http"
	"instagram/app/internals/services/reaction_story/usecase"
	"instagram/common"
)

func BuildReactStoryService(con common.SQLDatabase, v1 *gin.RouterGroup) {

	reactionRepo := reactstorymysql.NewMySQLStorage(con)
	getStoryRepo := rpc_client.NewGetStoryRepo(con)
	getReactionStory := reactionstoryusecase.NewInsertReactionStoryUseCase(reactionRepo, getStoryRepo)
	reactionStoryHandler := reactionstoryhttp.NewReactionStoryHandler(getReactionStory)
	reactionStoryHandler.RegisterV1Router(v1)
}
