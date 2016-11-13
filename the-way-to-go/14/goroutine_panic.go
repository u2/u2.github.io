package main

import (
	"fmt"
)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}

	defer close(ch)
}

func main() {
	ch := make(chan int)
	go tel(ch)
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	}
}
