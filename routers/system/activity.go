package system

import (
	v2 "github.com/forevengetmathmmqqmm/goGinExample/routers/api/v2"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct{}

func (u *ActivityRouter) InitActivityRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("activity")
	{
		apiRouter.POST("/add", v2.AddActivity)
		apiRouter.POST("/edit", v2.EditActivity)
		apiRouter.DELETE("/del/:id", v2.DelActivity)
		apiRouter.GET("/list", v2.ActivityList)
		apiRouter.GET("/detail/:id", v2.ActivityDetail)
	}
}
