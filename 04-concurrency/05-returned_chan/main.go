package main

import (
	"fmt"
	"time"
)

func main() {
	c := numberPump(0, 100*time.Millisecond) // sends every 100ms, starting at 0
	for {
		num := <-c
		fmt.Println(num)
	}
}

func numberPump(start int, sleep time.Duration) chan int {
	c := make(chan int)
	go func() {
		for i := start; ; i++ {
			c <- i
			time.Sleep(sleep)
		}
	}()
	return c
}
