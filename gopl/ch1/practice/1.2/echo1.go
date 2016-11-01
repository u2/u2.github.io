// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		t := strconv.Itoa(i)
		s += sep + t + " " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}
