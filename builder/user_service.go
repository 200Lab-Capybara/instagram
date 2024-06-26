package builder

import (
	"github.com/gin-gonic/gin"
	usermysql "instagram/app/infras/services/user/repository/mysql"
	userhttp "instagram/app/infras/services/user/transport/http"
	usersubscriber "instagram/app/infras/services/user/transport/subscriber"
	userusecase "instagram/app/internals/services/user/usecase"
	"instagram/components/pubsub/natspubsub"
	"instagram/components/tokenprovider"
)

func BuildUserService(serviceContext ServiceContext, accessPro tokenprovider.Provider, middleware gin.HandlerFunc) {
	con := serviceContext.GetDB()
	hasher := serviceContext.GetHasher()
	natsCon := serviceContext.GetNatsConn()
	v1 := serviceContext.GetV1()

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
