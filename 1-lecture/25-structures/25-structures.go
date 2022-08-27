package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 1}

	fmt.Println(p)
	fmt.Printf("%#+v\n", p)

	var p2 Point // zero value for structure is structure with zero values

	fmt.Println(p2)
	fmt.Printf("%#+v\n", p2)
}
