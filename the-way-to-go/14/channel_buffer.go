package main

import "fmt"

//import "time"

func main() {
	c := make(chan int, 50)
	c <- 10
	go func() {
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	fmt.Println("sent", 10)
}

/* Output:
sending 10
sent 10   // prints immediately
no further output, because main() then stops
*/
