package main

import (
	"sync"
	"sync/atomic"
	"time"
)

// Количество выполняемых горутин используя mutex
func MutexCounter() int {
	goroutinesCount := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			goroutinesCount++
			m.Unlock()

			time.Sleep(time.Microsecond)

			m.Lock()
			goroutinesCount--
			m.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCount
}

// Количество выполняемых горутин используя atomic
func AtomicCounter() int32 {
	goroutinesCounter := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goroutinesCounter, 1)
			time.Sleep(time.Microsecond)
			atomic.AddInt32(&goroutinesCounter, -1)

			wg.Done()
		}()
	}

	wg.Wait()
	return goroutinesCounter
}
