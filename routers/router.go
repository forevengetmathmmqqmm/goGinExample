package routers

import (
	"github.com/forevengetmathmmqqmm/goGinExample/middleware/jwt"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/setting"
	v2 "github.com/forevengetmathmmqqmm/goGinExample/routers/api/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiV2 := r.Group("/api/v2")
	apiV2.Use(jwt.JWT())
	{
		//获取标签列表
		apiV2.GET("/tags", v2.GetTags)
		apiV2.POST("/login", v2.Login)
		apiV2.GET("/test", v2.Test2)
	}
	return r
}
