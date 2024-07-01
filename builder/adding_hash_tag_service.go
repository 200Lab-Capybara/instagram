package builder

import (
	"github.com/gin-gonic/gin"
	hashtagsql "instagram/app/infras/services/hashtag/infras/repository/mysql"
	hashtaghttp "instagram/app/infras/services/hashtag/infras/transport/http"
	hashtagusercase "instagram/app/internals/services/hashtag/usecase"
	"instagram/common"
)

func BuildHashTagService(con common.SQLDatabase, v1 *gin.RouterGroup) {
	hashtagPostRepo := hashtagsql.NewMySQLStorage(con)
	hashtagCreatedRepo := hashtagsql.NewMySQLStorage(con)
	hashtagUC := hashtagusercase.NewCreateHashTagUseCase(hashtagPostRepo, hashtagCreatedRepo)
	hashtagHandler := hashtaghttp.NewHashTagHandler(hashtagUC)
	hashtagHandler.RegisterV1Router(v1)
}
