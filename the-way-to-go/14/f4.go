// Just cached
package main
import (
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	fmt.Println("11111th fibonacci is", Fibc(11109))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}

// subf processes two input channels in a wrapped goroutine
// and returns resulting new channels for subsequent fibonaccis
func subf(c1, c2 *int) (int, int) {
	*c1, *c2 = *c2, *c1 + *c2
	return *c1, *c2
}

// Fibc returns the nth fibonacci concurrently
func Fibc(n int) int {
	if n == 0 {
		return 0
	}

	a := 0
	b := 1

	// first calculation
	c, d := subf(&a, &b)

	// find desired nth fib
	for i := 1; i < n; i++ {
		c, d = subf(&c, &d)
	}
	return c
}
