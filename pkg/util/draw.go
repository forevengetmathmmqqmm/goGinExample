package util

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

func DrawImg(text string) (val string) {
	rand.Seed(time.Now().UnixNano())
	// 创建一个新的图像上下文
	width := 400
	height := 400
	dc := gg.NewContext(width, height)
	// 设置随机背景颜色
	bgR := rand.Float64()
	bgG := rand.Float64()
	bgB := rand.Float64()
	dc.SetRGB(bgR, bgG, bgB)
	dc.Clear()

	// 随机选择卡通图形颜色
	cartoonR := rand.Float64()
	cartoonG := rand.Float64()
	cartoonB := rand.Float64()
	dc.SetRGB(cartoonR, cartoonG, cartoonB)

	// 随机生成卡通图形位置和大小
	x := rand.Float64() * float64(dc.Width())
	y := rand.Float64() * float64(dc.Height())
	size := rand.Float64()*80 + 50
	dc.DrawEllipse(x, y, size, size*1.2)
	dc.Fill()
	cartoonR1 := rand.Float64()
	cartoonG1 := rand.Float64()
	cartoonB1 := rand.Float64()
	dc.SetRGB(cartoonR1, cartoonG1, cartoonB1)
	x1 := rand.Float64() * float64(dc.Width())
	y1 := rand.Float64() * float64(dc.Height())
	size1 := rand.Float64()*80 + 50
	dc.DrawEllipse(x1, y1, size1, size1*1.2)
	dc.Fill()
	cartoonR2 := rand.Float64()
	cartoonG2 := rand.Float64()
	cartoonB2 := rand.Float64()
	dc.SetRGB(cartoonR2, cartoonG2, cartoonB2)
	x2 := rand.Float64() * float64(dc.Width())
	y2 := rand.Float64() * float64(dc.Height())
	size2 := rand.Float64()*80 + 50
	dc.DrawEllipse(x2, y2, size2, size2*1.2)
	dc.Fill()
	cartoonR3 := rand.Float64()
	cartoonG3 := rand.Float64()
	cartoonB3 := rand.Float64()
	dc.SetRGB(cartoonR3, cartoonG3, cartoonB3)
	x3 := rand.Float64() * float64(dc.Width())
	y3 := rand.Float64() * float64(dc.Height())
	size3 := rand.Float64()*80 + 50
	dc.DrawEllipse(x3, y3, size3, size3*1.2)
	dc.Fill()
	cartoonR4 := rand.Float64()
	cartoonG4 := rand.Float64()
	cartoonB4 := rand.Float64()
	dc.SetRGB(cartoonR4, cartoonG4, cartoonB4)
	x4 := rand.Float64() * float64(dc.Width())
	y4 := rand.Float64() * float64(dc.Height())
	size4 := rand.Float64()*80 + 50
	dc.DrawEllipse(x4, y4, size4, size4*1.2)
	dc.Fill()
	dc.Fill()
	cartoonR5 := rand.Float64()
	cartoonG5 := rand.Float64()
	cartoonB5 := rand.Float64()
	dc.SetRGB(cartoonR5, cartoonG5, cartoonB5)
	x5 := rand.Float64() * float64(dc.Width())
	y5 := rand.Float64() * float64(dc.Height())
	size5 := rand.Float64()*80 + 50
	dc.DrawEllipse(x5, y5, size5, size5*1.2)
	dc.Fill()
	// 添加随机文本
	dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 24)
	dc.SetRGB(0, 0, 0)
	scaleFactor := 2.0
	dc.Scale(scaleFactor, scaleFactor)
	dc.DrawStringAnchored(text, 100, 100, 0.5, 0.5)
	outputFile, err := os.Create("uploads/random_cartoon.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	// 将图像保存为 PNG 文件
	err = dc.EncodePNG(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	return "/images/random_cartoon.png"
}
