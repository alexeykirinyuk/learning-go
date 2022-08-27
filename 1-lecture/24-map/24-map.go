package main

import "fmt"

func main() {
	var m map[string]int = make(map[string]int, 1) // у мапы нет капасити

	m["one"] = 1 // if not initialized - panic: assignment to entry in nil map

	fmt.Println(m, ", len(m) = ", len(m))
}
