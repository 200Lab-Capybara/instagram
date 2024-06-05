package userhttp

import (
	"github.com/nghiatrann0502/instagram-clone/common"
	"log"
	"net/http"
)

func (u *userHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Pong")
	common.SimpleSuccess(w, http.StatusOK, "pong")
}
