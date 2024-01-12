package system

import (
	v1 "github.com/forevengetmathmmqqmm/goGinExample/api/v1"
	"github.com/gin-gonic/gin"
)

type AccessRouter struct{}

func (a *AccessRouter) InitRoleRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("access")
	accessApi := v1.ApiGroupApp.SystemApiGroup.AccessApiGroup
	{
		apiRouter.POST("parent/add", accessApi.AddRole)
		apiRouter.DELETE("del/:id", accessApi.DelRoleInfoId)
		apiRouter.GET("list", accessApi.GetRoleList)
		apiRouter.GET("detail/:id", accessApi.GetRoleDetail)
	}
}
