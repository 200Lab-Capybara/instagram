package userhttp

import (
	"github.com/gin-gonic/gin"
	userusecase "instagram/app/internals/services/user/usecase"
)

type userHandler struct {
	registerUseCase userusecase.RegisterUseCase
}

func NewUserHandler(registerUseCase userusecase.RegisterUseCase) *userHandler {
	return &userHandler{
		registerUseCase: registerUseCase,
	}
}

func (u *userHandler) RegisterV1Router(v1 *gin.RouterGroup) {
	v1.POST("/register", u.GinRegisterHandler)
}
