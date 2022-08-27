package main

import (
	"fmt"

	"github.com/alexeykirinyuk/learning-go/2-workshop/storage/internal/storage"
)

func main() {
	fmt.Println("hello world")

	st := storage.NewStorage()
	fmt.Println(st)
}
