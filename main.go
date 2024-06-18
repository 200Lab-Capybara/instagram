package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	usermysql "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/repository/mysql"
	userhttp "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/transport/http"
	userusecase "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/usecase"
	"github.com/nghiatrann0502/instagram-clone/builder"
	"github.com/nghiatrann0502/instagram-clone/common"
	"github.com/nghiatrann0502/instagram-clone/components/hasher"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	v1 := r.Group("/v1")
	// Connect to database
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	bcrypt := hasher.NewBcryptHasher()
	con := common.NewSQLDatabase(db)

	// Create user storage
	userStorage, err := usermysql.NewMySQLStorage(con)
	if err != nil {
		log.Fatal(err)
	}

	registerUseCase := userusecase.NewRegisterUseCase(userStorage, bcrypt)
	userHandler := userhttp.NewUserHandler(registerUseCase)
	userHandler.RegisterV1Router(r)

	builder.BuildReactPostService(con, v1)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(httpAddr)
	if err != nil {
		log.Fatal(err)
	}
}
