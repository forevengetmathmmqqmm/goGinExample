package system

import (
	v1 "github.com/forevengetmathmmqqmm/goGinExample/api/v1"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (m *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("role")
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApiGroup
	{
		apiRouter.POST("edit", roleApi.AddRole)
		apiRouter.DELETE("del/:id", roleApi.DelRoleInfoId)
		apiRouter.GET("list", roleApi.GetRoleList)
		apiRouter.GET("detail/:id", roleApi.GetRoleDetail)
	}
}
