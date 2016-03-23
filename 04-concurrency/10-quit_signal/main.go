package main

import (
	"fmt"
	"time"
)

func main() {
	// how would we send a quit to multiple goroutines?
	quit := make(chan struct{})
	c := numberPump(0, 100, quit)
	timeout := time.After(2 * time.Second)

loop:
	for {
		select {
		case n := <-c:
			fmt.Println(n)
		case <-timeout:
			fmt.Println("timeout")
			quit <- struct{}{}
			break loop
		}
	}

	//time.Sleep(1 * time.Second) // uncomment to see quit message in numberPump
	fmt.Println("done.")
}

func numberPump(start int, sleep int, quit chan struct{}) chan int {
	// buffered channel. allow one entry in c before blocking on send, so we don't
	// get stuck sending to c while in main's timeout block.
	c := make(chan int, 1)
	go func() {
		i := start
		sleep := time.Duration(sleep) * time.Millisecond
		ticker := time.NewTicker(sleep)
		for {
			select {
			case <-ticker.C:
				c <- i
				i++
			case <-quit:
				fmt.Printf("quit: start:%d cur:%d sleep:%v\n", start, i, sleep)
				ticker.Stop()
				return
			}
		}
	}()
	return c
}
