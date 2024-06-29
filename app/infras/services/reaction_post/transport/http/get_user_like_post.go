package reactionposthttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

// POST /v1/posts/:id/like

func (hdl *reactionPostHandler) GetUserLikePostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid post ID"})
			return
		}

		result, err := hdl.getUserLikePostUC.Execute(c.Request.Context(), postId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
