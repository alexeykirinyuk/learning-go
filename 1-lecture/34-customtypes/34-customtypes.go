package main

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) SetX(newX float64) { // передача параметра по ЗНАЧЕНИЮ
	p.X = newX
}

func (p *Point) SetX2(newX float64) {
	p.X = newX
}

func main() {
	p := Point{1, 2}

	p.SetX(321)

	fmt.Println(p.X) // 1

	p.SetX2(100)
	fmt.Println(p.X) // 100
}
