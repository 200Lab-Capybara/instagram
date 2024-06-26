package reactionstoryhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

func (hdl *reactionStoryHandler) ReactStoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := common.User1UUID
		storyId, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		_, err = hdl.uc.Execute(c.Request.Context(), storyId, userId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
