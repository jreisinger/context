package main

import "fmt"

type Doubler struct {
	ctx *Context
	in  <-chan int
	out chan int
}

func NewDoubler(ctx *Context, in <-chan int) *Doubler {
	d := new(Doubler)
	d.in = in
	d.out = make(chan int)
	d.ctx = ctx

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := d.ctx.GetDone()
		for {
			select {

			// case d.out <- <-d.in * 2: // WRONG
			case i := <-d.in:
				select {
				case d.out <- i * 2:
				case <-done:
					fmt.Printf("Doubler terminated (2)\n")
				}

			case <-done:
				fmt.Printf("Doubler terminated\n")
				return
			}
		}
	}()

	return d
}

func (d *Doubler) GetSource() <-chan int {
	return d.out
}
