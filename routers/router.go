package routers

import (
	"github.com/forevengetmathmmqqmm/goGinExample/middleware/jwt"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/setting"
	"github.com/forevengetmathmmqqmm/goGinExample/routers/system"
	"github.com/gin-gonic/gin"
)

var apiV2 *gin.RouterGroup

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/images", "./uploads")
	gin.SetMode(setting.RunMode)
	apiV2 = r.Group("/api")
	apiV2.Use(jwt.JWT())
	system.Routers.InitUserRouter(apiV2)
	system.Routers.InitCommentRouter(apiV2)
	system.Routers.InitMusicianRouter(apiV2)
	system.Routers.InitActivityRouter(apiV2)
	system.Routers.InitBannerRouter(apiV2)
	system.Routers.InitRoleRouter(apiV2)
	system.Routers.InitAccessRouter(apiV2)
	return r
}
