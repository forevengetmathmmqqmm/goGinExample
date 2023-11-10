package v2

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/forevengetmathmmqqmm/goGinExample/models"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/app"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func AddActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	var req models.AddActivityRequest
	var res interface {
	}
	c.ShouldBind(&req)
	valid.Required(req.Name, "name").Message("活动名称不能为空")
	valid.Required(req.ImgUrl, "img_url").Message("活动封面不能为空")
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if valid.HasErrors() {
		code = e.ERROR
		msg = valid.Errors[0].String()
		appG.Response(http.StatusOK, code, msg, res)
	} else {
		sqlData, err := models.AddActivity(req)
		if err != nil {
			appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
		} else {
			res = sqlData
			appG.Response(http.StatusOK, code, msg, res)
		}
	}
}
func EditActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	var req models.EditActivityRequest
	var res interface {
	}
	c.ShouldBind(&req)
	valid.Required(req.Name, "name").Message("活动名称不能为空")
	valid.Required(req.ImgUrl, "img_url").Message("活动封面不能为空")
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if valid.HasErrors() {
		code = e.ERROR
		msg = valid.Errors[0].String()
		appG.Response(http.StatusOK, code, msg, res)
	} else {
		sqlData, err := models.EditActivity(req)
		if err != nil {
			appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
		} else {
			res = sqlData
			appG.Response(http.StatusOK, code, msg, res)
		}
	}
}
func ActivityList(c *gin.Context) {
	appG := app.Gin{C: c}
	res := new(models.ActivityListRequest)
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	sqlData, count, err := models.GetActivityList()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
	} else {
		res.List = sqlData
		res.Count = int(count)
		appG.Response(http.StatusOK, code, msg, res)
	}
}

func ActivityDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	sqlData, err := models.GetActivityDetail(id)
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
	} else {
		appG.Response(http.StatusOK, code, msg, sqlData)
	}
}

func DelActivity(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	err := models.DelActivity(id)
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
	} else {
		appG.Response(http.StatusOK, code, msg, nil)
	}
}
