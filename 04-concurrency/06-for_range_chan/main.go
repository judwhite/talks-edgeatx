package main

import (
	"fmt"
	"time"
)

func main() {
	c := numberPump(0, 100*time.Millisecond) // sends every 100ms, starting at 0
	for num := range c {                     // loop until channel closes
		fmt.Println(num)
	}
	fmt.Println("done.")
}

func numberPump(start int, sleep time.Duration) chan int {
	c := make(chan int)
	go func() {
		for i := start; i < start+10; i++ {
			c <- i
			time.Sleep(sleep)
		}
		close(c) // close chan, causes range to stop
	}()
	return c
}
