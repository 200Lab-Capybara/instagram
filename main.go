package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	usermysql "instagram/app/infras/services/user/repository/mysql"
	"instagram/builder"
	"instagram/common"
	"instagram/components/hasher"
	"instagram/components/tokenprovider"
	"instagram/middleware"
	"log"
	"net/http"
	"os"
)

var (
	httpAddr = os.Getenv("PORT")
	//":8001"
	connectionString = os.Getenv("CONNECTION_STRING")
	//"capybara:my_secret@tcp(localhost:3306)/users?parseTime=true"
)

func main() {
	r := gin.Default()
	r.Use(middleware.HandleError())
	v1 := r.Group("/v1")
	// Connect to database
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	bcrypt := hasher.NewBcryptHasher()
	con := common.NewSQLDatabase(db)

	accessTokenProvider := tokenprovider.NewJWTProvider(os.Getenv("ACCESS_SECRET"))
	//refreshTokenProvider := tokenprovider.NewJWTProvider(os.Getenv("REFRESH_SECRET"))

	userStorage := usermysql.NewMySQLStorage(con)
	authMiddleware := middleware.RequiredAuth(userStorage, accessTokenProvider)

	builder.BuildUserService(con, bcrypt, accessTokenProvider, v1, authMiddleware)
	builder.BuildReactPostService(con, v1)
	builder.BuildReactStoryService(con, v1)
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.HandleError())
	err = r.Run(httpAddr)
	if err != nil {
		log.Fatal(err)
	}
}
