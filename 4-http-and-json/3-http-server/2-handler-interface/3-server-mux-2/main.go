package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// так не подходит, нужен указатель
	// var _ http.Handler = &http.ServeMux{}

	// можно так
	mux := http.NewServeMux()

	mux.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "One")
	})

	mux.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Two")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
