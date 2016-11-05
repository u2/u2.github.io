package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c", "c", "dd", "dd", "e"}

	fmt.Println(a)
	a = uniq(a)

	fmt.Println(a)
}

func uniq(s []string) []string {
	n := 0
	for i, _ := range s {
		if i > 1 && s[i-1] == s[i] {
			// fmt.Println(s[i-1])
			// fmt.Println(s[i])
			copy(s[i-1:], s[i:])
			n++
		}
	}
	return s[:n+1]
}
