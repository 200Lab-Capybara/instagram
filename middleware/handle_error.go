package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"instagram/builder"
	"instagram/common"
)

func HandleError(ctxSvr builder.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					ctxSvr.GetLogger().Error(appErr.Error())
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				ctxSvr.GetLogger().Error(appErr.Error())
				c.AbortWithStatusJSON(appErr.StatusCode, errors.New("internal server error"))
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
