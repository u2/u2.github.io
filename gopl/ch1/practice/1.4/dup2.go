// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	files_names := make(map[string]string)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, files_names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files_names)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t %s\n", n, line, files_names[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, files_names map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		names := files_names[input.Text()]
		if !strings.Contains(names, f.Name()) {
			files_names[input.Text()] += " " + f.Name()
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
