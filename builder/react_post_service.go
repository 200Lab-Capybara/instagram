package builder

import (
	"github.com/gin-gonic/gin"
	reactionpostmysql "github.com/nghiatrann0502/instagram-clone/app/infras/services/reaction_post/repository/mysql"
	"github.com/nghiatrann0502/instagram-clone/app/infras/services/reaction_post/repository/rpc_client"
	reactionposthttp "github.com/nghiatrann0502/instagram-clone/app/infras/services/reaction_post/transport/http"
	reactionpostusecase "github.com/nghiatrann0502/instagram-clone/app/internals/services/reaction_post/usecase"
	"github.com/nghiatrann0502/instagram-clone/common"
)

func BuildReactPostService(con common.SQLDatabase, v1 *gin.RouterGroup) {
	reactionRepo := reactionpostmysql.NewMySQLStorage(con)
	getPostRepo := rpc_client.NewGetPostRepo(con)
	reactionUC := reactionpostusecase.NewLikePostUseCase(reactionRepo, getPostRepo)
	reactionHDL := reactionposthttp.NewReactionPostHandler(reactionUC)
	reactionHDL.RegisterV1Router(v1)
}
