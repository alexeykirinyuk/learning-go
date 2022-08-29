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
	mu := &sync.RWMutex{}
	mu.RLock()

	for i := 0; i < 1000; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			cs["касса_1"]++
		}(i)
	}

	for i := 0; i < 1000; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			cs["касса_2"]++
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
