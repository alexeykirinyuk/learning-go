package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(1) // максимальное кол-во используемых процессоров

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("gorouting %v\n", i)
		}()
		runtime.Gosched() // позволяет передать управледние другой горутине
	}

	time.Sleep(time.Second * 1)
}
