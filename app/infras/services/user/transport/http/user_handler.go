package userhttp

import (
	"github.com/gin-gonic/gin"
	userusecase "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/usecase"
)

type userHandler struct {
	registerUseCase userusecase.RegisterUseCase
}

func NewUserHandler(registerUseCase userusecase.RegisterUseCase) *userHandler {
	return &userHandler{
		registerUseCase: registerUseCase,
	}
}

func (u *userHandler) RegisterV1Router(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("/register", u.GinRegisterHandler)
}
