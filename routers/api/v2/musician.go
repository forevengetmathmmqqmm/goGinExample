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

func AddMusician(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	var req models.MusicianRequest
	var res interface {
	}
	c.ShouldBind(&req)
	valid.Required(req.StageName, "stage_name").Message("艺名不能为空")
	valid.Required(req.WeChat, "we_chat").Message("微信号不能为空")
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if valid.HasErrors() {
		code = e.ERROR
		msg = valid.Errors[0].String()
		appG.Response(http.StatusOK, code, msg, res)
	} else {
		sqlData, err := models.AddMusician(req)
		if err != nil {
			appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
		} else {
			res = sqlData
			appG.Response(http.StatusOK, code, msg, res)
		}
	}
}

func EditMusician(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	var req models.MusicianRequest
	var res interface {
	}
	c.ShouldBind(&req)
	valid.Required(req.StageName, "stage_name").Message("艺名不能为空")
	valid.Required(req.WeChat, "we_chat").Message("微信号不能为空")
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if valid.HasErrors() {
		code = e.ERROR
		msg = valid.Errors[0].String()
		appG.Response(http.StatusOK, code, msg, res)
	} else {
		sqlData, err := models.EditMusician(req)
		if err != nil {
			appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
		} else {
			res = sqlData
			appG.Response(http.StatusOK, code, msg, res)
		}
	}
}

func MusicianList(c *gin.Context) {
	appG := app.Gin{C: c}
	res := new(models.MusicianListRequest)
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	sqlData, count, err := models.GetMusicianList()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
	} else {
		res.List = sqlData
		res.Count = int(count)
		appG.Response(http.StatusOK, code, msg, res)
	}
}

func MusicianDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	sqlData, err := models.GetMusicianDetail(id)
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
	} else {
		appG.Response(http.StatusOK, code, msg, sqlData)
	}
}

func DelMusician(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	err := models.DelMusician(id)
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, err.Error(), nil)
	} else {
		appG.Response(http.StatusOK, code, msg, nil)
	}
}
