package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 1, Y: 1}
	fmt.Println(p) // 1, 1

	fill(p)
	fmt.Println(p) // 1, 1 -> значение копируется в функцию

	f := func() {
		fmt.Println("tmp")
	}

	call(f)
}

func fill(p Point) {
	p.X = 2
	p.Y = 2
}

func call(f func()) {
	f()
}

// func fill(p Point, len int) { // функции нельзя перегружать
// 	p.X = 2
// 	p.Y = 2
// }
