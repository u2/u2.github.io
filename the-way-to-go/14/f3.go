package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("28th fibonacci is", Fibc(26))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}

// subf processes two input channels in a wrapped goroutine
// and returns resulting new channels for subsequent fibonaccis
func subf(c1, c2 chan int) (chan int, chan int) {
	o1, o2 := make(chan int), make(chan int)
	go func(c1, c2 chan int) {
		a, b := <-c1, <-c2
		o1 <- b
		o2 <- a + b
	}(c1, c2)
	return o1, o2
}

// Fibc returns the nth fibonacci concurrently
func Fibc(n int) int {
	if n == 0 {
		return 0
	}

	a := make(chan int)
	b := make(chan int)

	// first calculation
	c, d := subf(a, b)

	// boom
	a <- 0
	b <- 1

	// find desired nth fib
	for i := 1; i < n; i++ {
		c, d = subf(c, d)
	}
	return <-c
}
