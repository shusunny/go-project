package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f, g := 1, 0
	//因为下面加法运算完后才返回f，且最终打印值为f()也就是f的值。所以如果我们想返回从0开始的数列，需要将g初始赋值为0
	return func() int {
		f, g = g, f+g
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}