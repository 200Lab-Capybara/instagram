package builder

import (
	"github.com/gin-gonic/gin"
	mysqlcomment "instagram/app/infras/services/comment/repository/mysql"
	rpc_clientcomment "instagram/app/infras/services/comment/repository/rpc_client"
	httpcomment "instagram/app/infras/services/comment/transport/http"
	usecasecomment "instagram/app/internals/services/comments/usecase"
	"instagram/common"
)

func BuildCommentService(con common.SQLDatabase, v1 *gin.RouterGroup, mid gin.HandlerFunc) {
	commentStorage := mysqlcomment.NewMySQLStorage(con)
	checkPost := rpc_clientcomment.NewGetPostRepo(con)
	createUseCase := usecasecomment.NewCommentUseCase(commentStorage, checkPost)
	commentHandler := httpcomment.NewCommentHandler(createUseCase)
	commentHandler.CommentV1Router(v1, mid)
}
