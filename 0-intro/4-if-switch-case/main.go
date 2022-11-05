package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// simple if
	{
		if rand.Int() > 10 {
			fmt.Println("cond is true")
		}
	}

	// if 2
	{
		if i := rand.Int(); i != 10 {
			fmt.Println("i is not 10!")
		}
	}

	// if 2 with err
	{
		bytes, err := os.ReadFile("test_file")
		if err != nil {
			// handle error
			_ = bytes // work with file
		}

		// f недоступна
	}

	// if else
	{
		i := rand.Int()
		if i > 10 {
			fmt.Println("i > 10")
		} else if i < 10 {
			fmt.Println("i < 10")
		} else {
			fmt.Println("i == 10")
		}
	}

	// switch case (break не нужен)
	{
		action := os.Args[0]
		switch action {
		case "ударить врага":
			fmt.Println("hit enemy")
		case "убежать":
			fmt.Println("run away")
		default:
			fmt.Println("эх")
		}
	}

	// switch case condition
	{
		i := rand.Int()
		switch {
		case i > 10:
			fmt.Println("i > 10")
		case i < 10:
			fmt.Println("i < 10")
		default:
			fmt.Println("i > 10")
		}
	}

	// fallthrough
	{
		discount := 0
		iam := "молодец"
		switch iam {
		case "супер молодец":
			discount += 10
			fallthrough
		case "молодец":
			discount += 10
			fallthrough
		default:
			discount += 5
		}

		// discount is 25%
	}
}
