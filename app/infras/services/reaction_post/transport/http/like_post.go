package reactionposthttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

// POST /v1/posts/:id/like

func (hdl *reactionPostHandler) ReactPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := common.User1UUID
		postId, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		_, err = hdl.uc.Execute(c.Request.Context(), userId, postId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
