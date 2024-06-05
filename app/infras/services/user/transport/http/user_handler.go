package userhttp

import (
	userinterface "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/interface"
	"net/http"
)

type userHandler struct {
	userUseCase userinterface.UserUseCase
}

func NewUserHandler(userUseCase userinterface.UserUseCase) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (u *userHandler) RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/register", u.RegisterHandler)
	mux.HandleFunc("/api/v1/ping", u.PingHandler)
}
