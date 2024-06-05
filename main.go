package main

import (
	usermysql "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/repository/mysql"
	userhttp "github.com/nghiatrann0502/instagram-clone/app/infras/services/user/transport/http"
	userusecase "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/usecase"
	"github.com/nghiatrann0502/instagram-clone/components/hasher"
	"log"
	"net/http"
)

var (
	httpAddr = ":8001"
)

func main() {
	mux := http.NewServeMux()
	bcrypt := hasher.NewBcryptHasher()

	userStorage, err := usermysql.NewMySQLStorage(nil)
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
