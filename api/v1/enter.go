package v1

import "github.com/forevengetmathmmqqmm/goGinExample/api/v1/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
