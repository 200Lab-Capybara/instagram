package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	usermysql "instagram/app/infras/services/user/repository/mysql"
	"instagram/builder"
	"instagram/common"
	"instagram/components/hasher"
	logruslogger "instagram/components/logger/logrus"
	"instagram/components/tokenprovider"
	"instagram/middleware"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	httpAddr = os.Getenv("PORT")
	//":8001"
	connectionString = os.Getenv("CONNECTION_STRING")
	//"capybara:my_secret@tcp(localhost:3306)/users?parseTime=true"
	natsConnectionString = os.Getenv("NATS_CONNECTION_STRING")
	//	nats://localhost:4222
)

func main() {
	r := gin.Default()
	logger := logruslogger.NewLogrusLogger()
	logger.Info("Starting the server.....")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/v1")
	v1Internal := r.Group("/internal/v1/rpc")
	// Connect to database
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	bcrypt := hasher.NewBcryptHasher()
	con := common.NewSQLDatabase(db)

	// NOTE: Connect to NATS
	natsCon, err := nats.Connect(natsConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	serviceContext := builder.NewServiceContext(con, bcrypt, natsCon, logger, v1, v1Internal)

	go builder.BuildRpcService(serviceContext)

	accessTokenProvider := tokenprovider.NewJWTProvider(os.Getenv("ACCESS_SECRET"))
	//refreshTokenProvider := tokenprovider.NewJWTProvider(os.Getenv("REFRESH_SECRET"))

	userStorage := usermysql.NewMySQLStorage(con)
	authMiddleware := middleware.RequiredAuth(userStorage, accessTokenProvider)


	builder.BuildUserService(serviceContext, accessTokenProvider, authMiddleware)
	builder.BuildReactPostService(con, natsCon, v1, authMiddleware)
	builder.BuildPostService(serviceContext, authMiddleware)
	builder.BuildReactStoryService(con, v1)
	builder.BuildStoryService(con, v1, natsCon, authMiddleware)
	builder.BuildReactCommentService(con, v1)
	builder.BuildProfileService(con, v1)
	builder.BuildFollowService(serviceContext, authMiddleware)

	// NOTE: This is a simple internal service route
	// NOTE: internal/v1/...

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.HandleError(serviceContext))
	err = r.Run(httpAddr)
	if err != nil {
		log.Fatal(err)
	}
}
