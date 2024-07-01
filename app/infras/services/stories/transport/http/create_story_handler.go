package storyhttp

import (
	"github.com/gin-gonic/gin"
	storiesmodel "instagram/app/internals/services/stories/model"
	"instagram/common"
	"net/http"
)

func (hdl *storyHandler) CreateStoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.RequesterKey).(common.Requester)
		storyDTO := storiesmodel.CreateStory{}

		if err := c.ShouldBindJSON(&storyDTO); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		id, err := hdl.storyUC.Execute(c.Request.Context(), requester, &storyDTO)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}
}
