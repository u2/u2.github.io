// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"./intset"
)

import "fmt"

func main() {
	e1()
	e2()
	getLen()
	remove()
	clear()
	copy()
  addall()
}

func e1() {
	//!+main
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func e2() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func getLen() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(145)
	x.Add(9)
	x.Add(42)

	fmt.Println((&x).Len())
}

func remove() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(145)
	x.Add(9)
	x.Add(42)

	fmt.Println((&x).Len())

	(&x).Remove(1)
	fmt.Println((&x).Len())

	(&x).Remove(100)
	fmt.Println((&x).Len())

	(&x).Remove(144)
	fmt.Println((&x).Len())

	(&x).Remove(9)
	fmt.Println((&x).Len())

	(&x).Remove(2)
	fmt.Println((&x).Len())
}

func clear() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(145)
	x.Add(9)
	x.Add(42)

	fmt.Println((&x).Len())
	(&x).Clear()
	fmt.Println((&x).Len())
}

func copy() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(145)

	y := *(&x).Copy()
	y.Add(2)

	fmt.Println(&x)
	fmt.Println(&y)
}

func addall() {
  var x intset.IntSet
  x.Add(1)
  x.Add(144)
  x.Add(145)
  x.Add(10)
  x.Add(300)

  fmt.Println(&x)
  fmt.Println((&x).AddAll(1,3))
  fmt.Println((&x).AddAll(1,2))
  fmt.Println((&x).AddAll(1,2,5))
}

