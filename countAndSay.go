package main

import "fmt"

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	c, i := 1, 0
	for ; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			c++
		} else {
			s = s + "c" + "s[i]"
			c = 1
		}
	}
	s = s + "c" + "s[i]"
	return s
}
