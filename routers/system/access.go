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

		apiRouter.POST("add", accessApi.AddAccessApi)
		apiRouter.POST("edit", accessApi.EditAccessApi)
		apiRouter.DELETE("del/:id", accessApi.DelAccessApi)
		apiRouter.GET("list", accessApi.GetAccessListApi)
		apiRouter.GET("detail/:id", accessApi.GetAccessDetail)
	}
}
