// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		t := strconv.Itoa(i)
		s += sep + t + " " + arg
		sep = "\n"
	}
	fmt.Println(s)
}
