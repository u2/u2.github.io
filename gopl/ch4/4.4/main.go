package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4}
	out := rotate(a, 3)
	fmt.Println(out)
}

func rotate(s []int, position int) []int {
	out := make([]int, len(s))
	for i, v := range s {
		if i < position {
			out[len(s)-position+i] = v
		} else {
			out[i-position] = v
		}
	}
	return out
}
