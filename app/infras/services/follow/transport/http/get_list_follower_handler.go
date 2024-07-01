package followhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
	"net/http"
)

func (hdl *followUserHandler) GetListFollowerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		paging := common.Paging{}
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		uid, err := uuid.Parse(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewCustomError(followusermodel.ErrInvalidUserId, followusermodel.ErrInvalidUserId.Error(), "invalid_user_id"))
			return
		}

		data, err := hdl.getFollowerUseCase.Execute(c.Request.Context(), uid, &paging)

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
