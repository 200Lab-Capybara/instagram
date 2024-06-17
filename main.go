package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	usermysql "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/repository/mysql"
	userhttp "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/transport/http"
	userusecase "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/usecase"
	"github.com/nghiatrann0502/instagram-clone/components/hasher"
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
)

func main() {
	r := gin.Default()
	// Connect to database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	bcrypt := hasher.NewBcryptHasher()

	userStorage, err := usermysql.NewMySQLStorage(db)
	if err != nil {
		log.Fatal(err)
	}
	userUseCase := userusecase.NewUserUseCase(userStorage, bcrypt)
	userHandler := userhttp.NewUserHandler(userUseCase)

	userHandler.RegisterGinHandler(r)

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
