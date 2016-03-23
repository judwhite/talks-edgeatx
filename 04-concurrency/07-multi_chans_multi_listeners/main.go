package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := numberPump(0, 100)
	c2 := numberPump(1000, 500)

	f := func(c chan int) {
		for num := range c {
			fmt.Println(num)
		}
	}

	go f(c1)
	go f(c2)

	//time.Sleep(2000 * time.Millisecond) // uncomment to allow goroutines to run for 2s
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
