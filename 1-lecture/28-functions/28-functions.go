package main

import "fmt"

func main() {
	fmt.Println(sum(5, 5, 5))

	slice := []int{5, 5}
	fmt.Println(sum(5, slice...)) // передача слайса в вариативные параметры
}

func sum(firstItem int, items ...int) (result int, ok bool) { // переменное кол-во элементов - items -> []int (слайс из int-ов)
	result = firstItem

	for _, val := range items {
		result += val
	}

	ok = true

	return

	// return 29, true // можно вернуть даже если именованные параметры
}
