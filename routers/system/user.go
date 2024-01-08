package system

import (
	v1 "github.com/forevengetmathmmqqmm/goGinExample/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApiGroup
	{
		apiRouter.POST("login", userApi.Login)
		apiRouter.POST("logout", userApi.Logout)
		apiRouter.POST("userInfo", userApi.EditUser)
		apiRouter.GET("userInfo/:id", userApi.GetUserDetail)
		apiRouter.GET("userInfo/list", userApi.GetUserList)
	}
}
