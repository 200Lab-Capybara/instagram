package postshttp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instagram/common"
	"net/http"
)

func (hdl *postHandler) GetListPostByUserIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postId, err := uuid.Parse(c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		paging := common.Paging{}

		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data, err := hdl.getPostsByUserId.Execute(c.Request.Context(), postId, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
