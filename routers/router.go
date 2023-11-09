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
	r.Static("/images", "./uploads")
	gin.SetMode(setting.RunMode)
	apiV2 := r.Group("/api/v2")
	apiV2.Use(jwt.JWT())
	{
		userApi := apiV2.Group("/user")
		{
			userApi.POST("/login", v2.Login)
			userApi.POST("/logout", v2.Logout)
			userApi.POST("/userInfo", v2.SetUserInfo)
			userApi.GET("/userInfo/:id", v2.GetUserDetail)
			userApi.GET("/userInfo/list", v2.GetUserList)
		}
		commentApi := apiV2.Group("/comment")
		{
			commentApi.POST("/uploads", v2.UploadFile)
			commentApi.GET("/draw", v2.DrawImg)
			commentApi.POST("/wallet", v2.UpdateWallet)
		}
		musicianApi := apiV2.Group("/musician")
		{
			musicianApi.POST("/add", v2.AddMusician)
			musicianApi.POST("/edit", v2.EditMusician)
			musicianApi.DELETE("/del/:id", v2.DelMusician)
			musicianApi.GET("/list", v2.MusicianList)
			musicianApi.GET("/detail/:id", v2.MusicianDetail)
		}
	}
	return r
}
