package hashtaghttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

// POST /v1/posts/:id/like
func (hdl *hashtagHandler) HashTagHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId, err := uuid.Parse(c.Param("Post_ID "))
		hashtagStrings := []string{"1", "2", "3"}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		_, err = hdl.uc.Execute(c.Request.Context(), postId, hashtagStrings)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
