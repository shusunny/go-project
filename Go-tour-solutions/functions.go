package main

import (
	"fmt"
	"math"
)

const delta = 1e-9 // 创建一个极小的常数

func Sqrt(x float64) float64 {
	z := x
	t := 0.0 // 临时变量，用来记录前一次z的值

	// 用math库中的绝对值函数计算两次的差值，当值差距较大时进入计算循环。
	for math.Abs(t-z) > delta {
		t, z = z, z-(z*z-x)/(2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("Our Result:\t", Sqrt(2))
	fmt.Println("Accurate Result:", math.Sqrt(2))
}