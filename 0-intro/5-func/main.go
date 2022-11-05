package main

import (
	"fmt"
	"math/rand"
)

func main() {
	printHelloRandom()

	printHelloForName("Алексей")

	printHello("Алексей", func(s string) int {
		fmt.Printf("Привет, дорогой %s\n", s)
		return 1
	})

	funcVariable := func(s string) int {
		fmt.Printf("Привет, дорогой %s\n", s)
		return 0
	}
	printHello("Алексей", funcVariable)
}

func printHelloRandom() {
	fmt.Printf("hello, person #%d\n", rand.Int())
}

//func printHelloRandom(newSignature int) {
//	fmt.Printf("hello, person #%d\n", rand.Int())
//}

func printHelloForName(name string) {
	fmt.Printf("hello, %s\n", name)
}

func printHello(name string, printFunc func(string) int) {
	printFunc(name)
}
