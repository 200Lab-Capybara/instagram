package httpcomment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	modelcomment "instagram/app/internals/services/comments/model"
	"instagram/common"
	"net/http"
)

func (hdl *commentHandler) CreateCommentHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := modelcomment.CommentCreation{}
		requester := c.MustGet(common.RequesterKey).(common.Requester)
		postId := c.Param("post_id")
		pId, err := uuid.Parse(postId)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		dto.PostId = pId

		_, err = hdl.createCommentUC.Execute(c.Request.Context(), requester, dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}
