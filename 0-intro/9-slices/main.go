package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("--- ARRAY, LEN")

	// arr, len
	{
		var arr [10]int
		for i := 0; i < len(arr); i++ {
			arr[i] = i
		}
		fmt.Printf("arr: %v. Len - %d, Cap - %d\n", arr, len(arr), cap(arr))
	}

	fmt.Println()

	// slice (range)
	fmt.Println("--- SLICE (range)")
	{
		arr := [10]int{}
		for i := 0; i < len(arr); i++ {
			arr[i] = i
		}

		sl := arr[0:3]
		sl2 := sl[0:1:1]
		_ = sl2
		fmt.Printf("slice: %v, len: %d, cap: %d\n", sl, len(sl), cap(sl))
		sl = append(sl, 100, 200)

		fmt.Printf("slice: %v, arr: %v\n", sl, arr)
	}

	fmt.Println()
	fmt.Println("--- SLICE")

	// slice
	{
		sl := make([]int64, 10) // (type, len, cap)
		for i := 0; i < 10; i++ {
			sl[i] = int64(i)
		}

		fmt.Println("sl: ", sl)
	}

	fmt.Println()
	fmt.Println("--- MAP 1")

	// map 1
	{
		m := map[string]int64{
			"one": 1,
			"two": 2,
		}
		fmt.Printf("%v\n", m)

		fmt.Println(m["one"])
	}

	fmt.Println()
	fmt.Println("--- MAP 2")

	// map 2
	{
		m := map[string][]string{
			"one": []string{"hello", "world"},
			"two": []string{"hello2", "world2"},
		}
		fmt.Printf("%v\n", m)

		fmt.Println(m["one"])
	}
}
