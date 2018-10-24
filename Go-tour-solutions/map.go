package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int) //创建map
	for _, c := range strings.Fields(s) {
		//遍历s，并对出现的每一个词汇在我们的计数器c上加1
		m[c]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}