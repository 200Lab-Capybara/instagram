package main

import (
	"database/sql"
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

	mux := http.NewServeMux()
	bcrypt := hasher.NewBcryptHasher()

	userStorage, err := usermysql.NewMySQLStorage(db)
	if err != nil {
		log.Fatal(err)
	}
	userUseCase := userusecase.NewUserUseCase(userStorage, bcrypt)
	userHandler := userhttp.NewUserHandler(userUseCase)
	userHandler.RegisterRouter(mux)

	log.Println("starting url server on", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
