package system

import (
	"time"

	"github.com/forevengetmathmmqqmm/goGinExample/models/common/response"
	"github.com/forevengetmathmmqqmm/goGinExample/models/system"
	server_system "github.com/forevengetmathmmqqmm/goGinExample/server/system"
	"github.com/forevengetmathmmqqmm/goGinExample/utils"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type AccessApiGroup struct{}

// 添加父级路由
func (a *AccessApiGroup) AddParentAccessApi(c *gin.Context) {
	var params system.AddParentAccess
	err := c.ShouldBind(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(params, utils.AddUserVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	params.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	user, err := server_system.AddParentAccess(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

// 编辑父级路由
func (a *AccessApiGroup) EditParentAccessApi(c *gin.Context) {
	var params system.EditParentAccess
	err := c.ShouldBind(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(params, utils.EditAccessVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	params.ModifiedOn = time.Now().Format("2006-01-02 15:04:05")
	accessParent, err := server_system.EditParentAccess(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(accessParent, c)
}

// 删除父级路由
func (a *AccessApiGroup) DelParentAccessApi(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0 {
		response.FailWithMessage("id为空", c)
	}
	err := server_system.DelAccessParentId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// 获取父级路由列表
func (a *AccessApiGroup) GetParentAccessListApi(c *gin.Context) {
	userList, count, err := server_system.GetAccessParentList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list":  userList,
		"total": count,
	}, c)
}

// 获取父级路由详情
func (u *AccessApiGroup) GetParentAccessDetail(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0 {
		response.FailWithMessage("id为空", c)
	}
	user, err := server_system.GetAccessParentDetail(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}
