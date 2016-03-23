package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var num int64
	go func() {
		for {
			atomic.AddInt64(&num, 1)
		}
	}()

	go func() {
		for {
			n := atomic.LoadInt64(&num)
			fmt.Println(n)
		}
	}()
	time.Sleep(1 * time.Second)
}
