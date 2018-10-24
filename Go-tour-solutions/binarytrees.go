package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
	//close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	if <-ch1 == <-ch2 {
		return true
	} else {
		return false
	}
}

func main() {
	ch := make(chan int)
	t1 := tree.New(1)
	t2 := tree.New(2)

	//从信道中打印10个值
	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	//对tree1和tree2进行比较
	fmt.Println("tree 1 == tree 1:", Same(t1, t1))
	fmt.Println("tree 1 == tree 2:", Same(t1, t2))
}
