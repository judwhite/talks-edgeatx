package main

import (
	"fmt"
	"time"
)

func main() {
	var num int

	go func() {
		for {
			num++
		}
	}()

	go func() {
		for {
			fmt.Println(num)
		}
	}()

	time.Sleep(1 * time.Second)
}
