package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	chArr := make([]chan int, 3)
	chArr[0] = make(chan int, 5)
	chArr[1] = make(chan int, 5)
	chArr[2] = make(chan int, 5)

	for i := 0; i < 3; i++ {
		go func(i int) {
			for y := 0; y < 5; y++ {
				chArr[2-i] <- y
				runtime.Gosched()
			}
		}(i)
	}

	go func() {
		for {
			select {
			case x := <-chArr[0]:
				fmt.Printf("<%v> -> %v\n", 0, x)
			case x := <-chArr[1]:
				fmt.Printf("<%v> -> %v\n", 1, x)
			case x := <-chArr[2]:
				fmt.Printf("<%v> -> %v\n", 2, x)
			}
		}
	}()
	time.Sleep(time.Second)
}
