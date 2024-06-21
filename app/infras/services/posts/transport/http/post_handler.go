package postshttp

import (
	"github.com/gin-gonic/gin"
	postusecase "instagram/app/internals/services/posts/usecase"
)

type postHandler struct {
	postUC postusecase.CreatePostUseCase
}

func NewPostHandler(postUC postusecase.CreatePostUseCase) *postHandler {
	return &postHandler{
		postUC: postUC,
	}
}

func (hdl *postHandler) RegisterV1Router(v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	v1.Use(middleware)
	v1.POST("/posts", hdl.CreatePostHandler())
}
