package main

import (
	"fmt"
	"time"
)

// имитатор wait group
// здесь куча проблем внутри
type wait struct {
	count uint
	ch    chan interface{}
}

func initWait(count uint) *wait {
	return &wait{
		count: count,
		ch:    make(chan interface{}, count), // значение интерфейса не важно, поэтому interface{}
	}
}

func (w *wait) Done() {
	w.ch <- 1
}

func (w *wait) Wait(id int) {
	// одна горутина может закрыть канал
	defer close(w.ch)

	for {
		<-w.ch

		// данная операция не атомарная
		// если Wait() будет вызван в разных горутинах - возможна ошибка
		w.count--

		if w.count <= 0 {
			return
		}
	}
}

func test(id int, w *wait) {
	defer w.Done()
	fmt.Printf("<#%d> done\n", id)
}

func main_() {
	w := initWait(5)

	for i := 0; i < 5; i++ {
		go test(i, w)
	}

	w.Wait(1)
	fmt.Println("done")
}

func main() {
	w := initWait(5)

	for i := 0; i < 5; i++ {
		go test(i, w)
	}

	go func() {
		w.Wait(2)
		fmt.Println("second done")
	}()

	w.Wait(1)
	fmt.Println("done")

	time.Sleep(time.Second)
}
