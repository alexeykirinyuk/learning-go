package main

import "fmt"

type PointWithComments struct {
	Point    Point
	Comments []string
}

type Point struct {
	X int
	Y int
}

func main() {
	var p PointWithComments
	fmt.Printf("%#+v\n", p)

	p2 := PointWithComments{Point: Point{X: 1, Y: 2}, Comments: []string{"comment 1", "comment 2"}}
	fmt.Printf("%#+v\n", p2)

	point1 := Point{X: 1, Y: 2}
	point2 := Point{X: 1, Y: 2}

	fmt.Println("point1 == point2 ->", point1 == point2)

	// Нельзя сравнить структуры на равенство, если ее поля можно сравнить
	// В данном случае нельзя сравнивать поле "Comments []string"
	// p3 := PointWithComments{Point: Point{X: 1, Y: 2}, Comments: []string{"comment 1", "comment 2"}}
	// fmt.Println("p2 == p3 ->", p2 == p3)
}
