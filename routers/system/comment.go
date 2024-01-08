package system

import (
	v2 "github.com/forevengetmathmmqqmm/goGinExample/routers/api/v2"
	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

func (u *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("comment")
	{
		apiRouter.POST("/uploads", v2.UploadFile)
		apiRouter.GET("/draw", v2.DrawImg)
		apiRouter.POST("/wallet", v2.UpdateWallet)
	}
}
