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

		if err := ctx.ShouldBindJSON(&dataCreation); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := dataCreation.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		id, err := u.registerUseCase.Execute(ctx.Request.Context(), &dataCreation)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}
}
