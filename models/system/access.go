package system

import (
	"github.com/forevengetmathmmqqmm/goGinExample/global"
)

type ParentAccess struct {
	global.Model
	Path        string `json:"path"`
	Icon        string `json:"icon"`
	Show        int    `json:"show"`
	Name        string `json:"name"`
	ElPath      string `json:"el_path"`
	HasChildren int    `json:"has_children"`
}

type Access struct {
	global.Model
	Path        string `json:"path"`
	Icon        string `json:"icon"`
	Show        int    `json:"show"`
	Name        string `json:"name"`
	ElPath      string `json:"el_path"`
	HasChildren int    `json:"has_children"`
}

type RoleAccess struct {
	global.Model
	AccessId int `json:"access_id"`
	RoleId   int `json:"role_id"`
}

// 编辑
type EditParentAccess struct {
	ID          int    `json:"id"`
	Path        string `json:"path"`
	Icon        string `json:"icon"`
	Show        int    `json:"show"`
	Name        string `json:"name"`
	ElPath      string `json:"el_path"`
	HasChildren int    `json:"has_children"`
	ModifiedOn  string `json:"modified_on"`
}

// 添加
type AddParentAccess struct {
	Path        string `json:"path"`
	Icon        string `json:"icon"`
	Show        int    `json:"show"`
	Name        string `json:"name"`
	ElPath      string `json:"el_path"`
	HasChildren int    `json:"has_children"`
	CreatedOn   string `json:"created_on"`
}
