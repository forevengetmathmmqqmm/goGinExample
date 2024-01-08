package system

import (
	v2 "github.com/forevengetmathmmqqmm/goGinExample/routers/api/v2"
	"github.com/gin-gonic/gin"
)

type MusicianRouter struct{}

func (m *MusicianRouter) InitMusicianRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("musician")
	// userApi := v1.ApiGroupApp.SystemApiGroup.UserApiGroup
	{
		apiRouter.POST("/add", v2.AddMusician)
		apiRouter.POST("/edit", v2.EditMusician)
		apiRouter.DELETE("/del/:id", v2.DelMusician)
		apiRouter.GET("/list", v2.MusicianList)
		apiRouter.GET("/detail/:id", v2.MusicianDetail)
	}
}
