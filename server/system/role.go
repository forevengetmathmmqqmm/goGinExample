package system

import (
	"github.com/forevengetmathmmqqmm/goGinExample/global"
	"github.com/forevengetmathmmqqmm/goGinExample/models/system"
)

func EditRole(params system.Role) (role system.Role, err error) {
	err = global.GAV_DB.Table("role").Save(&params).First(&role, params.ID).Error
	return role, err
}

// 获取角色列表
func GetRoleList() (role []system.Role, count int, err error) {
	err = global.GAV_DB.Table("role").Offset(0).Limit(10).Find(&role).Count(&count).Error
	return role, count, err
}
