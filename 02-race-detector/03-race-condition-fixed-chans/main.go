package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for num := 0; ; num++ {
			c <- num
		}
	}()

	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	time.Sleep(1 * time.Second)
}
