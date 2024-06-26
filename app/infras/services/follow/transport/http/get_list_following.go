package followhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

func (hdl *followUserHandler) GetFollowingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		paging := common.Paging{}
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		uid, err := uuid.Parse(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewCustomError(common.ErrInvalidRequest(err), "Invalid user id", "invalid_user_id"))
			return
		}
		data, err := hdl.getFollowingUseCase.Execute(c.Request.Context(), uid, &paging)

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
