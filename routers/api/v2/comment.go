package v2

import (
	"net/http"

	"github.com/forevengetmathmmqqmm/goGinExample/pkg/e"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/util"
	"github.com/gin-gonic/gin"
)

func DrawImg(c *gin.Context) {
	imgUrl := util.DrawImg()
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": imgUrl,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}
