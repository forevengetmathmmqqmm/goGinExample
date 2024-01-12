package system

import (
	"github.com/forevengetmathmmqqmm/goGinExample/global"
)

type ParentAccess struct {
	global.Model
	Path string `json:"path"`
	Icon string `json:"icon"`
	Show int    `json:"show"`
	Name string `json:"name"`
}

// 编辑
type EditParentAccess struct {
	ID         int    `json:"id"`
	Path       string `json:"path"`
	Icon       string `json:"icon"`
	Show       int    `json:"show"`
	Name       string `json:"name"`
	ModifiedOn string `json:"modified_on"`
}

// 添加
type AddParentAccess struct {
	Path      string `json:"path"`
	Icon      string `json:"icon"`
	Show      int    `json:"show"`
	Name      string `json:"name"`
	CreatedOn string `json:"created_on"`
}
