package main

import "fmt"

type EnrichedInt int // + можно объявить ссылочный тип *int например

func (t EnrichedInt) IsNegative() bool {
	return t < 0
}

func (t *EnrichedInt) Set(newValue int) {
	*t = EnrichedInt(newValue)
}

func main() {
	var i EnrichedInt = 4

	i -= 10

	var b int = 10

	// i -= b // нельзя так

	i -= EnrichedInt(b) // преобразование в EnrichedInt
	b -= int(i)         // преобразование в int

	fmt.Println(i.IsNegative())
	i.Set(-1)
	fmt.Println(i.IsNegative())
}
