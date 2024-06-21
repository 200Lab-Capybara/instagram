package httpprofile

import (
	"github.com/gin-gonic/gin"
	"instagram/app/internals/services/profile/model"
	"instagram/common"
	"net/http"
)

func (p *profileHandler) setupProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data model.ProfileCreation

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		id, err := p.setupUseCase.Execute(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}
}

func (p *profileHandler) SetupProfileV1Router(v1 *gin.RouterGroup) {
	v1.POST("/profile/setup", p.setupProfileHandler())
}
