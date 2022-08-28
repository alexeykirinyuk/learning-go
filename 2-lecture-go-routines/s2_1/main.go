package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func(x int, wg *sync.WaitGroup) { // нормальная практика передавать параметры в горутину с тем же названием
			defer wg.Done()
			fmt.Printf("gorouting #%d\n", x)
		}(x, &wg)
	}

	wg.Wait()
	fmt.Println("done")
}
