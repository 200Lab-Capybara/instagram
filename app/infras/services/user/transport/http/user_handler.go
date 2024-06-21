package userhttp

import (
	userusecase "instagram/app/internals/services/user/usecase"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	registerUseCase userusecase.RegisterUseCase
	loginUseCase    userusecase.LoginUseCase
}

func NewUserHandler(registerUseCase userusecase.RegisterUseCase, loginUC userusecase.LoginUseCase) *userHandler {
	return &userHandler{
		registerUseCase: registerUseCase,
		loginUseCase:    loginUC,
	}
}

func (u *userHandler) RegisterV1Router(v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	v1.POST("/register", u.RegisterHandler())
	v1.POST("/login", u.LoginHandler())

	requiredAuth := v1.Group("")
	requiredAuth.Use(middleware)
	requiredAuth.GET("/profile", middleware, u.GetProfileHandler())
}
