package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1))
	fmt.Println(max(1, 3, 5))
	fmt.Println(max(11, 24, 98))
	fmt.Println(max(901, 3))
}

func max(members ...int) int {
	t := 0
	for _, i := range members {
		if i > t {
			t = i
		}
	}
	return t
}
