package system

import (
	v1 "github.com/forevengetmathmmqqmm/goGinExample/api/v1"
	"github.com/gin-gonic/gin"
)

type AccessRouter struct{}

func (a *AccessRouter) InitAccessRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("access")
	accessApi := v1.ApiGroupApp.SystemApiGroup.AccessApiGroup
	{
		apiRouter.POST("parent/add", accessApi.AddParentAccessApi)
		apiRouter.POST("parent/edit", accessApi.EditParentAccessApi)
		apiRouter.DELETE("parent/del/:id", accessApi.DelParentAccessApi)
		apiRouter.GET("parent/list", accessApi.GetParentAccessListApi)
		apiRouter.GET("parent/detail/:id", accessApi.GetParentAccessDetail)

		apiRouter.POST("add", accessApi.AddParentAccessApi)
		apiRouter.POST("edit", accessApi.EditParentAccessApi)
		apiRouter.DELETE("del/:id", accessApi.DelParentAccessApi)
		apiRouter.GET("list", accessApi.GetParentAccessListApi)
		apiRouter.GET("detail/:id", accessApi.GetParentAccessDetail)
	}
}
