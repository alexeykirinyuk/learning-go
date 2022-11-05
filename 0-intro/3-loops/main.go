package main

import "fmt"

func main() {
	// классический цикл for
	{
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	// аналог white(true)
	{
		i := 0
		for {
			fmt.Println(i)

			if i > 10 {
				break
			}
			i++
		}
	}

	// опустить можно любой параметр
	{
		for i := 0; ; i++ {
			fmt.Println(i)

			if i > 10 {
				break
			}
		}
	}

	// аналог while (condition)
	{
		cond := true
		i := 0
		for cond {
			fmt.Println(i)
			i++
			cond = i > 10
		}
	}

	// foreach
	{
		strs := []string{"hello", "world"}
		for index, value := range strs {
			fmt.Println(index, value)
		}
	}
}
