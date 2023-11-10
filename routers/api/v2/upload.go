package v2

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/forevengetmathmmqqmm/goGinExample/pkg/app"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

type UploadRes struct {
	Path string `json:"path"`
}

func UploadFile(c *gin.Context) {
	appG := app.Gin{C: c}
	// 从请求中获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"message": "未上传文件"})
		return
	}
	uploadedFile, err := file.Open()
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, "文件无法打开", nil)
		return
	}
	defer uploadedFile.Close()
	img, _, err := image.Decode(uploadedFile)
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, "无法解码图像", nil)
		return
	}
	targetWidth := 200
	targetHeight := 200
	resizedImg := resize.Resize(uint(targetWidth), uint(targetHeight), img, resize.Lanczos3)
	// 将文件保存到本地
	outputFile, err := os.Create("uploads/" + file.Filename)
	err = jpeg.Encode(outputFile, resizedImg, nil)
	if err != nil {
		appG.Response(http.StatusInternalServerError, 500, "文件保存失败", nil)
		return
	}
	savedFilePath := "/images/" + file.Filename
	var res UploadRes
	res = UploadRes{
		Path: savedFilePath,
	}
	appG.Response(http.StatusOK, e.SUCCESS, e.GetMsg(e.SUCCESS), res)
}
