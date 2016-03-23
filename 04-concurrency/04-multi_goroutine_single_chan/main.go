package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go numberPump(0, 100*time.Millisecond, c)    // sends every 100ms, starting at 0
	go numberPump(1000, 500*time.Millisecond, c) // sends every 500ms, starting at 1000

	for {
		num := <-c
		fmt.Println(num)
	}
}

func numberPump(start int, sleep time.Duration, c chan int) {
	for i := start; ; i++ {
		c <- i
		time.Sleep(sleep)
	}
}
