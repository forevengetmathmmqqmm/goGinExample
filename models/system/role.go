package system

import "github.com/forevengetmathmmqqmm/goGinExample/global"

type Role struct {
	global.Model
	Title  string `json:"title" from:"title"`
	Desc   string `json:"desc" from:"desc"`
	Status int    `json:"status" from:"status"`
}
