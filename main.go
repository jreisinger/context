// context demonstrates how to terminate multiple goroutines. It's based on the
// ch. 8 of "Intermediate Go Programming" video course by John Graham-Cumming.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // for clean termination of the program

func main() {
	ctx := NewContext()

	c := NewCounter(ctx)
	connect := c.GetSource()

	d := NewDoubler(ctx, connect)
	read := d.GetSource()

	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)

	ctx.Stop()

	wg.Wait()
}
