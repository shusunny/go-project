package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)

	xs := strings.Split("Who is my faverate person, I don't know, the only thing I know is I don't like anyone")
	for
	}
}

func john(s []string) string{
	return strings.Join(s)
}
