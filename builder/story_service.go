package builder

import (
	"github.com/gin-gonic/gin"
	nats2 "github.com/nats-io/nats.go"
	storymysql "instagram/app/infras/services/stories/repository/mysql"
	storyhttp "instagram/app/infras/services/stories/transport/http"
	storysubscriber "instagram/app/infras/services/stories/transport/subscriber"
	storyusecase "instagram/app/internals/services/stories/usecase"
	"instagram/common"
	"instagram/components/pubsub/natspubsub"
)

func BuildStoryService(con common.SQLDatabase, v1 *gin.RouterGroup, pubsubCon *nats2.Conn, middleware gin.HandlerFunc) {
	storyStorage := storymysql.NewMysqlStorage(con)
	//reateStoryImages := storyrpcclient.NewCreateStoryImages(con)
	pubsub := natspubsub.NewNatsProvider(pubsubCon)
	createStoryUseCase := storyusecase.NewCreateStoryUC(storyStorage, pubsub)
	increaseLikeCountUC := storyusecase.NewIncreaseLikeCountUseCase(storyStorage)
	decreaseLikeCountUC := storyusecase.NewDecreaseLikeCountUseCase(storyStorage)

	// // HTTP Handler
	storyHandler := storyhttp.NewPostHandler(createStoryUseCase)
	storyHandler.RegisterV1Router(v1, middleware)

	// // Story Subscriber
	storySubscriber := storysubscriber.NewStorySubscriber(pubsub, increaseLikeCountUC, decreaseLikeCountUC)
	storySubscriber.Init()

}
