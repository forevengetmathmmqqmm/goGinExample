package main

import (
	setting "github.com/forevengetmathmmqqmm/goGinExample/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": setting.PageSize,
		})
	})
	r.Run()
}
