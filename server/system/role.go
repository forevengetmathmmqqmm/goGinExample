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
func GetRoleList() (role []system.Role, count int64, err error) {
	err = global.GAV_DB.Table("role").Offset(0).Limit(10).Find(&role).Count(&count).Error
	return role, count, err
}

func GetRole(id int) (role system.Role, err error) {
	err = global.GAV_DB.Table("role").First(&role, id).Error
	return role, err
}

type DelRoleResponse struct {
	Id int `json:"id"`
}

func DelRoleId(id int) (err error) {
	delData := DelRoleResponse{
		Id: id,
	}
	err = global.GAV_DB.Table("role").Delete(&delData).Error
	return err
}
