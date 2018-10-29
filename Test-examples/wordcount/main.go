package main

import (
	"./quote"
	"./word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	fmt.Println("------")

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, "times for word", k)
	}
}
