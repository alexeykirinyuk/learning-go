package main

import (
	"sync/atomic"
	"testing"
)

func Benchmark_AtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AtomicCounter()
	}
}

func Benchmark_MutexCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MutexCounter()
	}
}

func Benchmark_Inc(b *testing.B) {
	i1 := 0

	for i := 0; i < b.N; i++ {
		i1++
	}
}

func Benchmark_Atomic(b *testing.B) {
	i1 := int32(0)

	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&i1, 1)
	}
}

// results
// 1. Название функции (8 проц)
// 2. кол-во выполненных функций за время
// Benchmark_AtomicCounter-8           4076            276302 ns/op          104174 B/op       2003 allocs/op
// Benchmark_MutexCounter-8            3153            380056 ns/op          112353 B/op       2006 allocs/op
