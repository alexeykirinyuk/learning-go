package main

import "fmt"

func main() {

	// способ 1
	{
		var count int = 1
		fmt.Println(count)
	}

	// способ 2
	{
		var count = 1
		fmt.Println(count)
	}

	// способ 3. Самый распространенный. Эквивалентно способу 2
	{
		count := 1
		fmt.Println(count)
	}

	// zero-value
	{
		var count string
		fmt.Println(count)
	}

	// способ 1 может потребоваться для явного указания типа
	{
		var count int64 = 1
		fmt.Println(count)
	}

	// также можно привести к конкретному типу сразу
	{
		count := int64(1)
		fmt.Println(count)
	}
}
