package system

import (
	"time"

	"github.com/forevengetmathmmqqmm/goGinExample/models/common/response"
	"github.com/forevengetmathmmqqmm/goGinExample/models/system"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/util"
	server_system "github.com/forevengetmathmmqqmm/goGinExample/server/system"
	"github.com/forevengetmathmmqqmm/goGinExample/utils"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type UserApiGroup struct{}

// 添加
func (u *UserApiGroup) AddUserApi(c *gin.Context) {
	var user system.User
	err := c.ShouldBindJSON(*&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(user, utils.UserVerify)
}

// 登录
func (u *UserApiGroup) Login(c *gin.Context) {
	var params system.LoginApi
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(params, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := server_system.GetUserDetail(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if user.ID == 0 {
		response.FailWithMessage("账号或密码错误", c)
		return
	} else {
		token, _ := util.GenerateToken(params.Nickname, params.Password)
		loginData := system.LoginResp{
			Id:    user.ID,
			Token: token,
		}
		util.Blacklist[params.Nickname] = true
		response.OkWithDetailed(loginData, "登录成功", c)
	}
}

// 登出
func (u *UserApiGroup) Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	claims, err := util.ParseToken(token)
	if err != nil {
		response.FailWithMessage("Token鉴权失败", c)
	} else {
		util.Blacklist[claims.Nickname] = false
	}
	response.OkWithMessage("登出成功", c)
}

// 获取用户详情
func (u *UserApiGroup) GetUserDetail(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0 {
		response.FailWithMessage("用户id为空", c)
	}
	user, err := server_system.GetUserInfoId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

// 编辑用户信息
func (u *UserApiGroup) EditUser(c *gin.Context) {
	var params system.EditUserApi
	err := c.ShouldBind(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(params, utils.EditUserVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	params.ModifiedOn = time.Now().Format("2006-01-02 15:04:05")
	user, err := server_system.EditUser(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

// 获取用户列表
func (u *UserApiGroup) GetUserList(c *gin.Context) {
	userList, count, err := server_system.GetUserList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list":  userList,
		"total": count,
	}, c)
}
