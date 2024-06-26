package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
	"instagram/components/tokenprovider"
	"net/http"
	"strings"
)

var (
	ErrorUnauthorized = errors.New("invalid token")
)

type AuthStore interface {
	FindUserById(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
}

func RequiredAuth(authStore AuthStore, provider tokenprovider.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.GetHeader("Authorization")
		paths := strings.Split(author, " ")
		if len(paths) != 2 || paths[0] != "Bearer" || paths[1] == "" {
			c.JSON(http.StatusUnauthorized, common.NewUnauthorized(ErrorUnauthorized, ErrorUnauthorized.Error(), "invalid_token"))
			c.Abort()
			return
		}

		token := paths[1]

		claims, err := provider.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.NewUnauthorized(ErrorUnauthorized, ErrorUnauthorized.Error(), "invalid_token"))
			c.Abort()
			return
		}

		userId := claims.GetUserID()
		user, err := authStore.FindUserById(c.Request.Context(), userId)
		if err != nil {
			if errors.Is(err, usermodel.ErrUserNotFound) {
				c.JSON(http.StatusUnauthorized, common.NewUnauthorized(ErrorUnauthorized, ErrorUnauthorized.Error(), "invalid_token"))
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		if user.Status != common.UserActive {
			c.JSON(http.StatusUnauthorized, common.NewUnauthorized(ErrorUnauthorized, ErrorUnauthorized.Error(), "invalid_token"))
			c.Abort()
			return
		}

		requester := common.NewRequester(user.ID, user.FirstName, user.LastName, user.Role.String(), user.Status.String(), user.Follower, user.Following)

		c.Set(common.RequesterKey, requester)

		c.Next()
	}
}
