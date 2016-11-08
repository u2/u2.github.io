package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"a", "b"}
	b := []string{"1", "2"}
	fmt.Println(join(a))
	fmt.Println(join(a, b))
}

func join(members ...[]string) string {
	t := make([]string, 0)
	for _, i := range members {
		t = append(t, i...)
	}
	return strings.Join(t, ",")
}
