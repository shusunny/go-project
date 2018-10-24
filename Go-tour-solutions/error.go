package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} //引入错误处理的方法
	z := x
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	} //这里我只用了简单的10次重复计算，也可以用第一次练习中的任何牛顿法函数
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}