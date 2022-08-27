package main

import (
	"fmt"
	"log"

	"github.com/alexeykirinyuk/learning-go/2-workshop/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("main.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it's uploaded", file)

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it's loaded", restoredFile)
}
