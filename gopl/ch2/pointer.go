package main

import (
  "fmt"
)

func main() {
  f1()

  f2()


  fmt.Println(f() == f()) // "false"

  fmt.Println(*f() == *f()) // "false"
}

func f1() {
  x := 1
  p := &x         // p, of type *int, points to x
  fmt.Println(*p) // "1"
  *p = 2          // equivalent to x = 2
  fmt.Println(x)  // "2"
}

func f2() {
  var x, y int
  fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
}

func f() *int {
    v := 1
    return &v
}

