package system

import (
	"time"

	"github.com/forevengetmathmmqqmm/goGinExample/models/common/response"
	"github.com/forevengetmathmmqqmm/goGinExample/models/system"
	server_system "github.com/forevengetmathmmqqmm/goGinExample/server/system"
	"github.com/forevengetmathmmqqmm/goGinExample/utils"
	"github.com/gin-gonic/gin"
)

type RoleApiGroup struct{}

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
