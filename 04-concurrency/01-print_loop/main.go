package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; ; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}
