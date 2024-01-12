package system

import (
	"time"

	"github.com/forevengetmathmmqqmm/goGinExample/models/common/response"
	"github.com/forevengetmathmmqqmm/goGinExample/models/system"
	server_system "github.com/forevengetmathmmqqmm/goGinExample/server/system"
	"github.com/forevengetmathmmqqmm/goGinExample/utils"
	"github.com/gin-gonic/gin"
)

type AccessApiGroup struct{}

// 添加父级路由
func (a *AccessApiGroup) AddParentAccessApi(c *gin.Context) {
	var params system.AddParentAccess
	err := c.ShouldBind(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	err = utils.Verify(params, utils.AddUserVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	params.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	user, err := server_system.AddUser(params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}
