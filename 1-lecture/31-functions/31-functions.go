package main

import "fmt"

func main() {
	map1 := map[string]int{ // является "ссылочным" типом
		"one": 1,
	}
	fmt.Println(map1) // one:1

	fill(map1)
	fmt.Println(map1) // one:1, two:2
}

func fill(map1 map[string]int) {
	map1["two"] = 2
}
