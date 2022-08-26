package main

import "fmt"

func main() {
	c := 10
	slice := make([]int, 1, c) // (type, len, cap) or (type, len). if cap < len -> panic
	var slice2 []int           // nil слайс ведет себя также как и слайс из 0 элементов
	var slice3 []int = []int{}
	slice4 := slice[:1:1]

	slice[0] = 1000000

	printSlice(slice)
	printSlice(slice2)
	printSlice(slice4)
	fmt.Println("slice2 == nil", slice2 == nil)
	fmt.Println("slice3 == nil", slice3 == nil)
}

func printSlice(slice []int) {
	fmt.Println(slice)
	fmt.Println("let(slice)", len(slice))
	fmt.Println("cap(slice)", cap(slice))
	fmt.Println("===")
	fmt.Println()
}
