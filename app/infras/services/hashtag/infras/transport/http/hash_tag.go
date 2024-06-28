package hashtaghttp

import (
	"github.com/gin-gonic/gin"
	"instagram/common"
	"net/http"
)

// POST /v1/posts/:id/like
func (hdl *hashtagHandler) HashTagHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := common.Post1UUID
		hashtagStrings := []string{"4", "9"}

		success, err := hdl.uc.Execute(c.Request.Context(), postId, hashtagStrings)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if !success {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to react to the post"})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
