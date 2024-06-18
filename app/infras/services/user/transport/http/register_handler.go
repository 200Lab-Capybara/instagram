package userhttp

import (
	"github.com/gin-gonic/gin"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
	"net/http"
)

func (u *userHandler) RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		id, err := u.registerUseCase.Execute(ctx.Request.Context(), &dataCreation)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}
}
