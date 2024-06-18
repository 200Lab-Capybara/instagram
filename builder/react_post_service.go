package builder

import (
	"github.com/gin-gonic/gin"
	reactionpostmysql "instagram/app/infras/services/reaction_post/repository/mysql"
	"instagram/app/infras/services/reaction_post/repository/rpc_client"
	reactionposthttp "instagram/app/infras/services/reaction_post/transport/http"
	reactionpostusecase "instagram/app/internals/services/reaction_post/usecase"
	"instagram/common"
)

func BuildReactPostService(con common.SQLDatabase, v1 *gin.RouterGroup) {
	reactionRepo := reactionpostmysql.NewMySQLStorage(con)
	getPostRepo := rpc_client.NewGetPostRepo(con)
	reactionUC := reactionpostusecase.NewLikePostUseCase(reactionRepo, getPostRepo)
	reactionHDL := reactionposthttp.NewReactionPostHandler(reactionUC)
	reactionHDL.RegisterV1Router(v1)
}
