package builder

import (
	"github.com/gin-gonic/gin"
	hashtagsql "instagram/app/infras/services/hashtag/infras/repository"
	userhttp "instagram/app/infras/services/user/transport/http"
	hashtagusercase "instagram/app/internals/services/hashtag/usecase"
	"instagram/common"
)

func BuildAddingHashTagService(con common.SQLDatabase, v1 *gin.RouterGroup) {
	hashtagRepo := hashtagsql.NewMySQLStorage(con)
	hashtagUC := hashtagusercase.NewCreateHashTagUseCase(hashtagRepo)
	hashtagHandler := userhttp.NewUserHandler(hashtagUC)
	hashtagHandler.RegisterV1Router(v1)
	//hashtagHDL := reactionposthttp.NewCreateHashTagHandler(hashtagUC)
	//hashtagHDL.RegisterV1Router(v1)
}
