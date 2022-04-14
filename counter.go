package main

import "fmt"

type Counter struct {
	ctx *Context
	c   chan int
	// done chan struct{}
	i int
}

func NewCounter(ctx *Context) *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.ctx = ctx

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := counter.ctx.GetDone()
		for {
			select {
			case counter.c <- counter.i:
				counter.i++
			case <-done:
				fmt.Printf("Counter terminated\n")
				return
			}
		}
	}()

	return counter
}

func (c *Counter) GetSource() <-chan int {
	return c.c
}
