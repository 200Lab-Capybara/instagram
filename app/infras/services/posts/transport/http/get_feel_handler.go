package postshttp

import (
	"github.com/gin-gonic/gin"
	"instagram/common"
	"net/http"
)

func (hdl *postHandler) GetFeedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user id from the path
		requester := c.MustGet(common.RequesterKey).(common.Requester)

		paging := common.Paging{}
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Call the use case
		posts, err := hdl.getFeedUC.Execute(c.Request.Context(), requester, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(posts, paging, nil))
	}
}
