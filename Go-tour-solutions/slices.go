package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var ret = make([][]uint8, dy)
	//外层切片
	for x := 0; x < dy; x++ {
		ret[x] = make([]uint8, dx)
		//内层切片
		for y := 0; y < dx; y++ {
			ret[x][y] = uint8((x + y) / 2)
			//给内层切片每个元素赋值，可以更改为题目中的其他函数
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}