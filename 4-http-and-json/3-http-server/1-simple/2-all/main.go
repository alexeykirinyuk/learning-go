package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello, world!")
	}

	http.HandleFunc("/", helloHandler)

	err := http.ListenAndServe(":5001", nil) // сервер доступен снаружи
	if err != nil {
		log.Fatal(err)
	}
}
