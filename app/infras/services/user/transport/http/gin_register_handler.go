package userhttp

import (
	"github.com/gin-gonic/gin"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
	"net/http"
)

func (u *userHandler) GinRegisterHandler(ctx *gin.Context) {
	dataCreation := usermodel.UserCreation{}

	err := ctx.ShouldBindJSON(&dataCreation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	if err := dataCreation.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := u.userUseCase.Register(ctx.Request.Context(), &dataCreation)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": id,
	})
}
