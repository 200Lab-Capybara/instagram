package builder

import (
	"github.com/gin-gonic/gin"
	postsmysql "instagram/app/infras/services/posts/repository/mysql"
	postrpcclient "instagram/app/infras/services/posts/repository/rpc_client"
	postshttp "instagram/app/infras/services/posts/transport/http"
	postusecase "instagram/app/internals/services/posts/usecase"
	"instagram/common"
)

func BuildPostService(con common.SQLDatabase, v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	postStorage := postsmysql.NewMysqlStorage(con)
	createPostImages := postrpcclient.NewCreatePostImages(con)
	postUC := postusecase.NewCreatePostUseCase(postStorage, createPostImages)
	postHandler := postshttp.NewPostHandler(postUC)
	postHandler.RegisterV1Router(v1, middleware)
}
