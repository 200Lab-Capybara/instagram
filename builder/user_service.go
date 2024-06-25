package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	usermysql "instagram/app/infras/services/user/repository/mysql"
	userhttp "instagram/app/infras/services/user/transport/http"
	usersubscriber "instagram/app/infras/services/user/transport/subscriber"
	userusecase "instagram/app/internals/services/user/usecase"
	"instagram/common"
	"instagram/components/hasher"
	"instagram/components/pubsub/natspubsub"
	"instagram/components/tokenprovider"
)

func BuildUserService(con common.SQLDatabase, natsCon *nats.Conn, hasher hasher.Hasher, accessPro tokenprovider.Provider, v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	userStorage := usermysql.NewMySQLStorage(con)
	pubsub := natspubsub.NewNatsProvider(natsCon)

	// Usecase
	registerUseCase := userusecase.NewRegisterUseCase(userStorage, hasher)
	loginUC := userusecase.NewLoginUseCase(userStorage, hasher, accessPro)

	followUC := userusecase.NewUpdateFollowUseCase(userStorage)
	unfollowUC := userusecase.NewUpdateUnfollowUseCase(userStorage)

	userHandler := userhttp.NewUserHandler(registerUseCase, loginUC)
	userHandler.RegisterV1Router(v1, middleware)

	userSub := usersubscriber.NewUserSubscriber(pubsub, unfollowUC, followUC)
	userSub.Init()
}
