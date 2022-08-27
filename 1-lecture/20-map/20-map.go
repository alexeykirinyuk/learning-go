package main

import "fmt"

func main() {
	strToNum := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// local variables in some scope - example
	{
		a := 4
		fmt.Println(a)
	}

	if value, ok := strToNum["zero"]; ok { // value & ok определены для всей конструкции if/else, но не вне
		fmt.Println("Zero is inside map and it's value is", value)
	} else {
		fmt.Println("There are no zero element")
	}

	if value, ok := strToNum["zero-2"]; ok {
		fmt.Println("Zero-2 is inside map and it's value is", value)
	} else {
		fmt.Println("There are no zero-2 element")
	}
}

// different variable scopes
