package postshttp

import (
	"github.com/gin-gonic/gin"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
	"net/http"
)

func (hdl *postHandler) CreatePostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.RequesterKey).(common.Requester)
		postDTO := postsmodel.PostCreation{}

		if err := c.ShouldBindJSON(&postDTO); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		id, err := hdl.postUC.Execute(c.Request.Context(), requester, &postDTO)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}
}
