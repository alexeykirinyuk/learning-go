package main

import "fmt"

func main() {
	// ptr = <array>
	// len = 2
	// cap = 5
	array := [5]float64{1, 2, 3, 4, 5}

	// берем 0 и 1 индексы. Второй не берем (не включая правую границу)
	slice := array[:2]
	print(slice)

	slice = append(slice, 100)
	print(slice)
}

func print(slice []float64) {
	fmt.Println(slice)
	fmt.Println("let(slice)", len(slice))
	fmt.Println("cap(slice)", cap(slice))
	fmt.Println("===")
	fmt.Println()
}
