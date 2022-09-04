package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	i := 0

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello, world!")
		log.Printf("hello world %d", i)
		i++
	}

	http.HandleFunc("/", helloHandler)

	http.ListenAndServe("localhost:8080", nil) // блокирующая функция
}
