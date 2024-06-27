package builder

import (
	"github.com/gin-gonic/gin"
	mysqlreactcomment "instagram/app/infras/services/reaction_comment/repository/mysql"
	rpc_clientreactioncomment "instagram/app/infras/services/reaction_comment/repository/rpc_client"
	httpreactioncomment "instagram/app/infras/services/reaction_comment/transport/http"
	usecasereactioncomment "instagram/app/internals/services/reaction_comment/usecase"
	"instagram/common"
)

func BuildReactCommentService(con common.SQLDatabase, v1 *gin.RouterGroup) {
	reactionRepo := mysqlreactcomment.NewMySQLStorage(con)
	getCommentRepo := rpc_clientreactioncomment.NewGetCommentRepo(con)
	getReactionComment := usecasereactioncomment.NewInsertReactionCommentUseCase(reactionRepo, getCommentRepo)
	reactionCommentHandler := httpreactioncomment.NewReactionCommentHandler(getReactionComment)
	reactionCommentHandler.RegisterV1Router(v1)
}
