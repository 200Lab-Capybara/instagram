package main

import (
<<<<<<< HEAD
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
=======
	"github.com/gin-gonic/gin"
>>>>>>> main
	usermysql "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/repository/mysql"
	userhttp "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/transport/http"
	userusecase "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/usecase"
	"github.com/nghiatrann0502/instagram-clone/common"
	"github.com/nghiatrann0502/instagram-clone/components/hasher"
<<<<<<< HEAD
=======
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
>>>>>>> main
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

<<<<<<< HEAD
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		// log.Fatalf("failed to start server: %v", err)
=======
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(httpAddr)
	if err != nil {
		log.Fatal(err)
>>>>>>> main
	}
}
