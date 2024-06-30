package userhttp

import (
	"github.com/gin-gonic/gin"
	"instagram/common"
	"net/http"
)

func (hdl *userHandler) GetProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.RequesterKey).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(requester.GetSimpleUser()))
	}
}
