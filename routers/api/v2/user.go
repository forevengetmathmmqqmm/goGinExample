package v2

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/forevengetmathmmqqmm/goGinExample/models"
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
			} else {
				code = e.PASSWORD_FAIL
				msg = e.GetMsg(e.PASSWORD_FAIL)
			}
		} else {
			sqlData := models.Login(req)
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
func UserInfo(c *gin.Context) {

}
func GetUserDetail(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	sqlData := models.HasUser("id", id)
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
