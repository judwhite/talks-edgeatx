package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := numberPump(0, 100)
	c2 := numberPump(1000, 500)

	timeout := time.After(2 * time.Second)

loop:
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
		case <-timeout:
			fmt.Println("timeout")
			break loop
		}
	}
	fmt.Println("done.")
}

func numberPump(start int, sleep int) chan int {
	c := make(chan int)
	go func() {
		for i := start; ; i++ {
			c <- i
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return c
}
