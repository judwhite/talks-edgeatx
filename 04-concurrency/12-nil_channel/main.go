package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	quit := make(chan struct{})
	c1 := numberPump(0, 50, quit)     // time reduced to 50ms
	c2 := numberPump(1000, 300, quit) // time reduced to 300ms
	timeout := time.After(3 * time.Second)

	wg.Add(2)

loop:
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
			if n == 5 {
				//c1 = nil // uncomment to stop receiving on c1
			}
		case n := <-c2:
			fmt.Println(n)
		case <-timeout:
			fmt.Println("timeout")
			close(quit)
			break loop
		}
	}

	wg.Wait() // wait for all goroutines to call wg.Done()
	fmt.Println("done.")
}

func numberPump(start int, sleep int, quit <-chan struct{}) <-chan int {
	c := make(chan int) // no longer a buffered channel

	go func() {
		i := start
		ticker := time.NewTicker(time.Duration(sleep) * time.Millisecond)
		var msg string

	loop:
		for {
			select {
			case <-ticker.C:
				select {
				case c <- i:
					i++
				case <-quit:
					msg = "quit received in inner select"
					break loop
				}
			case <-quit:
				msg = "quit received in outer select"
				break loop
			}
		}

		fmt.Printf("quit: msg:%q start:%-4d cur:%-4d sleep:%v\n", msg, start, i, sleep)
		ticker.Stop()
		wg.Done()
	}()

	return c
}
