package util

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/fogleman/gg"
)

func DrawImg() (val string) {
	// 打开原始图片文件
	inputImageFile, err := os.Open("uploads/4.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer inputImageFile.Close()

	// 使用 `image` 包解码图像文件
	inputImage, _, err := image.Decode(inputImageFile)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个新的图像上下文
	width := inputImage.Bounds().Max.X
	height := inputImage.Bounds().Max.Y
	dc := gg.NewContext(width, height)

	// 将原始图片绘制到新的图像上下文中
	dc.DrawImage(inputImage, 0, 0)

	// 使用卡通化算法进行图像处理
	cartoonize(dc.Image())

	// 保存生成的卡通图片
	outputFile, err := os.Create("uploads/cartoonized.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// 将图像保存为 PNG 文件
	err = gg.SavePNG("uploads/cartoonized.png", dc.Image())
	if err != nil {
		log.Fatal(err)
	}
	return "/images/cartoonized.png"
}
func cartoonize(img image.Image) {
	// 在这里，我们使用简化的卡通化算法来模拟卡通风格
	rgba := image.NewRGBA(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			brightness := (r + g + b) / 3
			newColor := color.RGBA{
				R: uint8(brightness >> 8),
				G: uint8(brightness >> 8),
				B: uint8(brightness >> 8),
				A: 255,
			}
			rgba.Set(x, y, newColor)
		}
	}
}
