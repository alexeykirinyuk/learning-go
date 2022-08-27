package main

import "fmt"

func main() {
	strToNum := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}

	value, ok := strToNum["zero"]
	fmt.Println(value, ok) // 0 true

	value, ok = strToNum["not-found"]
	fmt.Println(value, ok) // 0 false
}
