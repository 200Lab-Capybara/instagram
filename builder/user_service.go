package builder

import (
	"github.com/gin-gonic/gin"
	usermysql "instagram/app/infras/services/user/repository/mysql"
	userhttp "instagram/app/infras/services/user/transport/http"
	userusecase "instagram/app/internals/services/user/usecase"
	"instagram/common"
	"instagram/components/hasher"
	"instagram/components/tokenprovider"
)

func BuildUserService(con common.SQLDatabase, hasher hasher.Hasher, accessPro tokenprovider.Provider, v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	userStorage := usermysql.NewMySQLStorage(con)

	// Usecase
	registerUseCase := userusecase.NewRegisterUseCase(userStorage, hasher)
	loginUC := userusecase.NewLoginUseCase(userStorage, hasher, accessPro)

	userHandler := userhttp.NewUserHandler(registerUseCase, loginUC)
	userHandler.RegisterV1Router(v1, middleware)
}
