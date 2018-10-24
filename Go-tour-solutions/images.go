package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

//ColorModel 应当返回 color.RGBAModel
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

//Bounds 应当返回一个 image.Rectangle ，例如 image.Rect(0, 0, w, h)
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

//At 应当返回一个颜色。
//上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}
func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2) //运用了和上次一样的函数
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}