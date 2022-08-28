package main

import (
	"fmt"
	"time"
)

// Пул тасок
type Pool struct {
	in      chan int
	out     chan string
	runners uint
	task    func(int) string
}

func NewPool(buffer uint) *Pool {
	return &Pool{
		in:      make(chan int),
		out:     make(chan string, buffer),
		runners: buffer,
	}
}

func (p *Pool) SetTask(f func(int) string) {
	p.task = f
}

func (p *Pool) Run() {
	counter := make(chan interface{}, p.runners)

	for {
		// получаем входные данные для таски
		in, ok := <-p.in
		if !ok {
			return
		}

		x := true
		for x {
			select {
			// если получилось записать в каунтер - значит есть свободные таски (занимаем ее)
			case counter <- 0:
				x = false
			// для дебага, увидим что свободных тасок в пуле нет
			default:
				fmt.Printf("%v\t-\n", in)
				time.Sleep(time.Millisecond * 100)
			}
		}
		go func(in int) {
			// при выполнении задачи освобождаем 1 таску из пула
			defer func() {
				<-counter
			}()

			// выполняем таску и отдаем в output
			p.out <- p.task(in)
		}(in)
	}
}

func (p *Pool) OutChan() <-chan string {
	return p.out
}

func (p *Pool) Send(i int) {
	p.in <- i
}

func main() {
	pool := NewPool(10)
	pool.SetTask(func(s int) string {
		t := "low"
		if s%2 == 0 {
			time.Sleep(time.Second)
		} else {
			t = "fast"
			time.Sleep(time.Millisecond * 300)
		}

		time.Sleep(time.Second)
		return fmt.Sprintf("%v\t%v", s, t)
	})

	go func() {
		for out := range pool.OutChan() {
			fmt.Println(out)
		}
	}()

	go func() {

		for i := 0; i < 100; i++ {
			go pool.Send(i)
		}
	}()

	go pool.Run()

	time.Sleep(time.Second * 10)
}
