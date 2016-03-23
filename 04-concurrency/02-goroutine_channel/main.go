package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)					// make a channel for communicating ints

	go numberPump(c)   					// start goroutine

	for {
		num := <-c 						// receive value on channel c, assign to num
		fmt.Println(num)
	}
}

func numberPump(c chan int) {
	for i := 0; ; i++ {
		c <- i 							// send value of i on channel c
		time.Sleep(100 * time.Millisecond)
	}
}