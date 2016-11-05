package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string]int)
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    word := scanner.Text()
    counts[word]++
  }

  if scanner.Err() != nil {
    fmt.Fprintln(os.Stderr, scanner.Err())
    os.Exit(1)
  }

  for w, n := range counts {
    fmt.Printf("%s\t%d\n", w, n)
  }
}
