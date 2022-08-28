package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	example5()
}

// сигнал завершения
func example5() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	generate := func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- 1
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}

	dataChan := make(chan int)

	go generate(dataChan)

	for {
		select {
		case <-c:
			fmt.Println("Получен сигнал завершения")
			return
		case data, ok := <-dataChan:
			if !ok {
				fmt.Println("dataChan закрылся")
				return
			}

			fmt.Println(data)
		}
	}
}

// wait группы (ожидание завршения потоков)
func example4() {
	longOperation := func(i int, wg *sync.WaitGroup) {
		defer wg.Done()

		time.Sleep(time.Second)
		fmt.Printf("%v passed\n", i)
	}

	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go longOperation(i, wg)
	}

	wg.Wait()

	fmt.Println("done")
}

// пример deadlock-а
func example3() {
	var ch chan int

	ch <- 1
	_, _ = <-ch
}

// ограничение по записи и чтении
// 1. закрывает канал тот кто в него пишет
// тот кто читает из канала - закрыть его не может (ошибка компиляции)
// 4. если пишет несколько продюсеров - закрывает канал тот, кто его создал
// 3. не закрытый канал держит ресурсы, закрывать нужно их явно
func example2() {
	readChannel := func(ch <-chan int) {
		for x := range ch {
			fmt.Println(x)
		}
	}

	writeChannel := func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second / 10)
		}
		close(ch)
	}

	ch := make(chan int, 5)
	go readChannel(ch)
	writeChannel(ch)
}

// простой пример с каналами
func example1() {
	var ch chan int

	fmt.Printf("cap(ch)=%d,len(ch)=%d\n", cap(ch), len(ch)) // cap(ch)=0,len(ch)=0

	ch = make(chan int, 5)
	fmt.Printf("cap(ch)=%d,len(ch)=%d\n", cap(ch), len(ch)) // cap(ch)=5,len(ch)=0

	ch <- 1
	fmt.Printf("cap(ch)=%d,len(ch)=%d\n", cap(ch), len(ch)) // cap(ch)=5,len(ch)=1
}
