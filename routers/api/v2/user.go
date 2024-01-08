package v2

import (
	"net/http"

	"github.com/forevengetmathmmqqmm/goGinExample/models"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/e"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/util"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func UpdateWallet(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	claims, err := util.ParseToken(token)
	var code = e.SUCCESS
	if err != nil {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	sqlData := models.HasUser("nickname", claims.Nickname)
	var respData models.WalletResp
	c.ShouldBind(&respData)
	models.UpdateWallet(respData, sqlData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
