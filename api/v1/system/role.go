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

type RoleApiGroup struct{}

// 获取角色详情
func (u *RoleApiGroup) GetRoleDetail(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0 {
		response.FailWithMessage("角色id为空", c)
	}
	user, err := server_system.GetRole(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

func (r *RoleApiGroup) DelRoleInfoId(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0 {
		response.FailWithMessage("角色id为空", c)
	}
	err := server_system.DelRoleId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// 编辑添加角色
func (r *RoleApiGroup) AddRole(c *gin.Context) {
	var params system.Role
	err := c.ShouldBind(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(params, utils.EditRoleVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if params.ID != 0 {
		params.ModifiedOn = time.Now().Format("2006-01-02 15:04:05")
	} else {
		params.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	}
	user, err := server_system.EditRole(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

// 角色列表
func (r *RoleApiGroup) GetRoleList(c *gin.Context) {
	userList, count, err := server_system.GetRoleList()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list":  userList,
		"total": count,
	}, c)
}
