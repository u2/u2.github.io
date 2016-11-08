package main

import (
	"fmt"
	"gopl.io/ch6/geometry"
)

func main() {
	perim := geometry.Path{geometry.Point{1, 1}, geometry.Point{5, 1}, geometry.Point{5, 4}, geometry.Point{1, 1}}
	fmt.Println(perim.Distance())
}
