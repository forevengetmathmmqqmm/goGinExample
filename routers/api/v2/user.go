package v2

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/forevengetmathmmqqmm/goGinExample/models"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/app"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/e"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type Res struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	valid := validation.Validation{}
	var req models.Req
	c.ShouldBind(&req)
	valid.Required(req.Nickname, "nickname").Message("昵称不能为空")
	valid.Required(req.Password, "password").Message("密码不能为空")
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	var data Res
	if valid.HasErrors() {
		code = e.ERROR
		msg = valid.Errors[0].String()
	} else {
		token, _ := util.GenerateToken(req.Nickname, req.Password)
		sqlData := models.HasUser("nickname", req.Nickname)
		if sqlData.ID > 0 {
			if req.Password == sqlData.Password {
				data = Res{
					sqlData.ID,
					token,
				}
				util.Blacklist[sqlData.Nickname] = true
			} else {
				code = e.PASSWORD_FAIL
				msg = e.GetMsg(e.PASSWORD_FAIL)
			}
		} else {
			sqlData := models.Login(req)
			util.Blacklist[sqlData.Nickname] = true
			data = Res{
				sqlData.ID,
				token,
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func SetUserInfo(c *gin.Context) {
	valid := validation.Validation{}
	var respData models.RespUser
	c.ShouldBind(&respData)
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	valid.Required(respData.ID, "id").Message("用户id不能为空")
	valid.Required(respData.Nickname, "nickname").Message("昵称不能为空")
	var resData models.ResUser
	if valid.HasErrors() {
		code = e.ERROR
		msg = valid.Errors[0].String()
	} else {
		resData = models.UpdateUser(respData)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": resData,
	})
}

func GetUserDetail(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	sqlData := models.GetUser("id", id)
	var code = e.SUCCESS
	if id < 0 {
		code = e.USER_FAIL
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": sqlData,
	})
}

// 获取用户列表
func GetUserList(c *gin.Context) {
	appG := app.Gin{C: c}
	sqlData := models.GetUserList()
	appG.Response(http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS), map[string]interface{}{
		"list":  sqlData,
		"total": 10,
	})
}
func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	var code = e.SUCCESS
	claims, err := util.ParseToken(token)
	if err != nil {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else {
		util.Blacklist[claims.Nickname] = false
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
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
