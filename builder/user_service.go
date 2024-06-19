package builder

import (
	"github.com/gin-gonic/gin"
	usermysql "instagram/app/infras/services/user/repository/mysql"
	userhttp "instagram/app/infras/services/user/transport/http"
	userusecase "instagram/app/internals/services/user/usecase"
	"instagram/common"
	"instagram/components/hasher"
)

func BuildUserService(con common.SQLDatabase, hasher hasher.Hasher, v1 *gin.RouterGroup) {
	userStorage := usermysql.NewMySQLStorage(con)
	registerUseCase := userusecase.NewRegisterUseCase(userStorage, hasher)
	userHandler := userhttp.NewUserHandler(registerUseCase)
	userHandler.RegisterV1Router(v1)
}
