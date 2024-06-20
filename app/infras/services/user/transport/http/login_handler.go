package userhttp

import (
	"github.com/gin-gonic/gin"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
	"net/http"
)

func (u *userHandler) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		dataLogin := usermodel.UserLogin{}

		if err := c.ShouldBindJSON(&dataLogin); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := dataLogin.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data, err := u.loginUseCase.Execute(c.Request.Context(), &dataLogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
