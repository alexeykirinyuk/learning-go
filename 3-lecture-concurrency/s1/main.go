package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{
		"касса_1": 0,
		"касса_2": 0,
	}
	mu1 := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func(i int) {
			mu1.Lock()
			defer mu1.Unlock()
			cs["касса_1"]++
		}(i)
	}

	mu2 := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			mu2.Lock()
			defer mu2.Unlock()
			cs["касса_2"]++
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
