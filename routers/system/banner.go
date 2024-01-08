package system

import (
	v2 "github.com/forevengetmathmmqqmm/goGinExample/routers/api/v2"
	"github.com/gin-gonic/gin"
)

type BannerRouter struct{}

func (u *BannerRouter) InitBannerRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("banner")
	{
		apiRouter.POST("/add", v2.AddBanner)
		apiRouter.POST("/edit", v2.EditActivity)
		apiRouter.DELETE("/del/:id", v2.DelActivity)
		apiRouter.GET("/list", v2.ActivityList)
		apiRouter.GET("/detail/:id", v2.ActivityDetail)
	}
}
