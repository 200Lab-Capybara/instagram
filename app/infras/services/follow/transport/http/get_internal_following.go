package followhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

func (hdl *followUserHandler) GetInternalFollowingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := struct {
			UserId uuid.UUID `json:"id"`
		}{}

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data, err := hdl.getInternalFollowingUC.Execute(c.Request.Context(), &dto.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
