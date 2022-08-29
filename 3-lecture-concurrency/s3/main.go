package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// m := sync.Mutex{}
	// m.Lock()

	value := 0
	value++
	fmt.Println(value)

	value2 := int32(0)
	atomic.AddInt32(&value2, 1)
	fmt.Println(value2)
}
