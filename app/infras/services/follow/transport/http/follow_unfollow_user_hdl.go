package followhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
	"net/http"
)

func (hdl *followUserHandler) FollowUnfollowUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.RequesterKey).(common.Requester)
		// Get the following id from the path
		followingID, err := uuid.Parse(c.Param("following_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewCustomError(followusermodel.ErrInvalidFollowingId, followusermodel.ErrInvalidFollowingId.Error(), "invalid_following_id"))
			return
		}

		// Call the use case
		ok, err := hdl.followUserUseCase.Execute(c.Request.Context(), requester, followingID)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(ok))
	}
}
