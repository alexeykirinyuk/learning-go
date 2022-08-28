package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("gorouting %v\n", i)
		}()
	}

	time.Sleep(time.Second * 1)
}
