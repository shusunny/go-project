package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// 创建一个外层长为dy的二维切片
	p := make([][]uint8, dy)
	// 遍历p, 为每一个外层的切片创建一个内层切片
	for x := range p {
		p[x] = make([]uint8, dx)
		for y := range p[x] {
			// 给每一个元素赋值，下面的函数可以更改为题目中的其他函数
			p[x][y] = uint8((x + y) / 2)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
