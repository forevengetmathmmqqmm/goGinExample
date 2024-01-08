package system

import (
	"github.com/forevengetmathmmqqmm/goGinExample/global"
	"github.com/forevengetmathmmqqmm/goGinExample/models/system"
)

func GetUserDetail(params system.LoginApi) (user system.User, err error) {
	err = global.GAV_DB.Table("user").Where("nickname = ? AND password >= ?", params.Nickname, params.Password).First(&user).Error
	return user, err
}

func GetUserInfoId(id int) (user system.User, err error) {
	err = global.GAV_DB.Table("user").First(&user, id).Error
	return user, err
}

func EditUser(params system.EditUserApi) (user system.User, err error) {
	err = global.GAV_DB.Table("user").Save(&params).First(&user, params.ID).Error
	return user, err
}

// 获取用户列表
func GetUserList() (userList []system.UserListResp, count int, err error) {
	err = global.GAV_DB.Table("user").Offset(0).Limit(10).Find(&userList).Count(&count).Error
	return userList, count, err
}
