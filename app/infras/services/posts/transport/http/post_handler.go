package postshttp

import (
	"github.com/gin-gonic/gin"
	postusecase "instagram/app/internals/services/posts/usecase"
)

type postHandler struct {
	createPostUseCase postusecase.CreatePostUseCase
	getPostsByUserId  postusecase.GetListPostByUserIdUseCase
	getFeedUC         postusecase.GetFeedUseCase
}

func NewPostHandler(postUC postusecase.CreatePostUseCase, getPostsByUserId postusecase.GetListPostByUserIdUseCase, getFeedUC postusecase.GetFeedUseCase) *postHandler {
	return &postHandler{
		createPostUseCase: postUC,
		getPostsByUserId:  getPostsByUserId,
		getFeedUC:         getFeedUC,
	}
}

func (hdl *postHandler) RegisterV1Router(v1 *gin.RouterGroup, middleware gin.HandlerFunc) {
	v1.POST("/posts", middleware, hdl.CreatePostHandler())
	v1.GET("/posts", middleware, hdl.GetFeedHandler())
	v1.GET("/posts/:userId/get-by-user", middleware, hdl.GetListPostByUserIdHandler())
}
