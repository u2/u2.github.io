package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "foo abc ddd foo bar xx hello world foo"
	s = expand(s, f)
	fmt.Println(s)
}

func f(s string) string {
	return strings.Replace(s, "foo", "bar", -1)
}

func expand(s string, f func(string) string) string {
	return f(s)
}
