package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1) // tons of numbers appear
	time.Sleep(15e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		time.Sleep(1e9)
		ch <- i
		fmt.Println(i)
	}
}

func suck(ch chan int) {
	time.Sleep(10e9)
	for {
		fmt.Println(<-ch)
	}
}
