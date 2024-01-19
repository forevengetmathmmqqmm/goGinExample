package system

import (
	"github.com/forevengetmathmmqqmm/goGinExample/global"
	"github.com/forevengetmathmmqqmm/goGinExample/models/system"
)

func AddParentAccess(params system.AddParentAccess) (parentAccess system.ParentAccess, err error) {
	err = global.GAV_DB.Table("access_parent").Create(&params).Find(&parentAccess, "path = ?", params.Path).Error
	return parentAccess, err
}

func EditParentAccess(params system.EditParentAccess) (parentAccess system.ParentAccess, err error) {
	err = global.GAV_DB.Table("access_parent").Save(&params).Find(&parentAccess, "path = ?", params.Path).Error
	return parentAccess, err
}

type DelAccessParentResponse struct {
	Id int `json:"id"`
}

func DelAccessParentId(id int) (err error) {
	delData := DelAccessParentResponse{
		Id: id,
	}
	err = global.GAV_DB.Table("access_parent").Delete(&delData).Error
	return err
}

// 获取列表
func GetAccessParentList() (list []system.ParentAccess, count int, err error) {
	err = global.GAV_DB.Table("access_parent").Offset(0).Limit(10).Find(&list).Count(&count).Error
	return list, count, err
}

// 获取详情
func GetAccessParentDetail(id int) (parentAccess system.ParentAccess, err error) {
	err = global.GAV_DB.Table("access_parent").First(&parentAccess, id).Error
	return parentAccess, err
}
