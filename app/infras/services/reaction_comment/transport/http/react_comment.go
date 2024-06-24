package httpreactioncomment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

func (hdl *reactionCommentHandler) ReactCommentHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := common.User1UUID
		commentId, err := uuid.Parse(c.Param("commentId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		_, err = hdl.uc.Execute(c.Request.Context(), userId, commentId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
