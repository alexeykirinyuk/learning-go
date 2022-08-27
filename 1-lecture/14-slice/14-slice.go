package main

import "fmt"

func main() {
	fmt.Println("description | first element pointer | len | cap | slice")

	a := []int{}
	print("before", a)

	a = append(a, 1)
	print("after 1", a)

	a = append(a, 2)
	print("after 2", a)

	a = append(a, 3)
	print("after 3", a)

	a = append(a, 4)
	print("after 4", a)

	a = append(a, 5)
	print("after 5", a)
}

func print(desc string, slice []int) {
	var pointer *int
	if len(slice) > 0 {
		pointer = &slice[0]
	}

	fmt.Printf(
		"%11s | %21p | %3d | %3d | %v \n",
		desc,
		pointer,
		len(slice),
		cap(slice),
		slice)
}
