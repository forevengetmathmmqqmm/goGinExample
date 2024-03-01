package system

import (
	"fmt"
	"sort"

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
func GetAccessParentList() (list []system.ParentAccess, count int64, err error) {
	err = global.GAV_DB.Table("access_parent").Offset(0).Limit(10).Find(&list).Count(&count).Error
	return list, count, err
}

// 获取详情
func GetAccessParentDetail(id int) (parentAccess system.ParentAccess, err error) {
	err = global.GAV_DB.Table("access_parent").First(&parentAccess, id).Error
	return parentAccess, err
}

func AddAccess(params system.AddAccess) (access system.Access, err error) {
	err = global.GAV_DB.Table("access").Create(&params).Find(&access, "path = ?", params.Path).Error
	return access, err
}

func AddAccessRole(params []system.RoleAccess) (err error) {
	err = global.GAV_DB.Table("role_access").Create(&params).Error
	return err
}

// 编辑role_access
func EditAccessRole(params []system.RoleAccess, accessId int) (err error) {
	global.GAV_DB.Table("role_access").Where("access_id = ?", accessId).Delete(&system.RoleAccess{})
	err = AddAccessRole(params)
	return err
}

type AccessList []system.Access

func (a AccessList) Len() int {
	return len(a)
}
func (p AccessList) Less(i, j int) bool {
	return p[i].ParentRouterId < p[j].ParentRouterId
}
func (p AccessList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 筛选role_access
func AccessRoleList(role_id int) (routers []system.Routers, err error) {
	roleAccessList := []system.RoleAccess{}
	global.GAV_DB.Table("role_access").Where("role_id = ?", role_id).Find(&roleAccessList)
	accessList := AccessList{}
	for _, v := range roleAccessList {
		access, _ := GetAccessDetail(v.AccessId)
		if access.Show == 1 {
			accessList = append(accessList, access)
		}
	}
	sort.Sort(accessList)
	middleList := []system.Access{}
	for i, v := range accessList {
		if i == len(accessList)-1 || v.ParentRouterId != accessList[i+1].ParentRouterId {
			middleList = append(middleList, v)
			parentAccessData, _ := GetAccessParentDetail(v.ParentRouterId)
			router := system.Routers{}
			router.ChildrenList = middleList
			router.ElPath = parentAccessData.ElPath
			router.HasChildren = parentAccessData.HasChildren
			router.ID = parentAccessData.ID
			router.Icon = parentAccessData.Icon
			router.Name = parentAccessData.Name
			router.Path = parentAccessData.Path
			routers = append(routers, router)
			middleList = []system.Access{}
		} else {
			middleList = append(middleList, v)
		}
	}
	fmt.Print("=====routers======", &routers)
	return routers, err
}

// 编辑路由
func EditAccess(params system.EditAccessParams) (access system.Access, err error) {
	err = global.GAV_DB.Table("access").Save(&params).Find(&access, "path = ?", params.Path).Error
	return access, err
}

// 获取列表
func GetAccessList() (list []system.Access, count int64, err error) {
	err = global.GAV_DB.Table("access").Offset(0).Limit(10).Find(&list).Count(&count).Error
	for i, v := range list {
		roleAccess := []system.RoleAccess{}
		global.GAV_DB.Table("role_access").Find(&roleAccess, "access_id = ?", v.ID)
		for _, a := range roleAccess {
			list[i].RoleIds = append(list[i].RoleIds, a.RoleId)
		}
	}
	return list, count, err
}

// 获取详情
func GetAccessDetail(id int) (access system.Access, err error) {
	err = global.GAV_DB.Table("access").First(&access, id).Error
	roleAccess := []system.RoleAccess{}
	global.GAV_DB.Table("role_access").Find(&roleAccess, "access_id = ?", id)
	for _, a := range roleAccess {
		access.RoleIds = append(access.RoleIds, a.RoleId)
	}
	return access, err
}

func DelAccessId(id int) (err error) {
	delData := DelAccessParentResponse{
		Id: id,
	}
	err = global.GAV_DB.Table("access").Delete(&delData).Error
	return err
}
