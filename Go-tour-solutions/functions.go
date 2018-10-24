package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	temp := 0.0 // 临时变量，用来记录前一次z的值
	for {
		z = z - (z*z-x)/(2*z) // 计算出最新的z值
		fmt.Println(z)
		if math.Abs(z-temp) < 0.000000000000001 {
			break // 用math库中的绝对值使值停止改变（或改变非常小）的时候退出循环
		} else {
			temp = z // 改变仍较大时记录此时的z值
		}
	}
	return z
}

func main() {
	fmt.Println("Our Result:\t", Sqrt(2))
	fmt.Println("Accurate Result:", math.Sqrt(2))
}