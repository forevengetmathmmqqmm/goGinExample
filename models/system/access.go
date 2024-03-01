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

type Routers struct {
	ParentAccess
	ChildrenList []Access `json:"children_list"`
}

type RoleAccess struct {
	AccessId   int    `json:"access_id"`
	RoleId     int    `json:"role_id"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
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

type Access struct {
	global.Model
	Path           string `json:"path"`
	Icon           string `json:"icon"`
	Show           int    `json:"show"`
	Name           string `json:"name"`
	ElPath         string `json:"el_path"`
	RoleIds        []int  `json:"role_ids"`
	ParentRouterId int    `json:"parent_router_id"`
}

// 编辑
type EditAccessParams struct {
	EditAccess
	RoleIds []int `json:"role_ids"`
}

// 编辑
type EditAccess struct {
	ID             int    `json:"id"`
	Path           string `json:"path"`
	Icon           string `json:"icon"`
	Show           int    `json:"show"`
	Name           string `json:"name"`
	ElPath         string `json:"el_path"`
	ModifiedOn     string `json:"modified_on"`
	ParentRouterId int    `json:"parent_router_id"`
}

// 添加
type AddAccess struct {
	Path           string `json:"path"`
	Icon           string `json:"icon"`
	Show           int    `json:"show"`
	Name           string `json:"name"`
	ElPath         string `json:"el_path"`
	ParentRouterId int    `json:"parent_router_id"`
	CreatedOn      string `json:"created_on"`
}

// 添加
type AddAccessParams struct {
	AddAccess
	RoleIds []int `json:"role_ids"`
}
