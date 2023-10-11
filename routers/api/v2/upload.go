package v2

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/forevengetmathmmqqmm/goGinExample/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

type UploadRes struct {
	Path string `json:"path"`
}

func UploadFile(c *gin.Context) {
	// 从请求中获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"message": "未上传文件"})
		return
	}
	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "文件无法打开"})
		return
	}
	defer uploadedFile.Close()
	img, _, err := image.Decode(uploadedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "无法解码图像"})
		return
	}
	targetWidth := 200
	targetHeight := 200
	resizedImg := resize.Resize(uint(targetWidth), uint(targetHeight), img, resize.Lanczos3)
	// 将文件保存到本地
	outputFile, err := os.Create("uploads/" + file.Filename)
	err = jpeg.Encode(outputFile, resizedImg, nil)
	if err != nil {
		c.JSON(500, gin.H{"message": "文件保存失败"})
		return
	}
	savedFilePath := "images/" + file.Filename
	var res UploadRes
	res = UploadRes{
		Path: savedFilePath,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "文件上传成功",
		"data": res,
	})
}
