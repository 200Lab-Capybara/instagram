package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"instagram/common"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, errors.New("internal server error"))
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
