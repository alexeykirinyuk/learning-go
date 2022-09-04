package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%+v\n", r)

		// Можно без форматирования строки
		_, _ = fmt.Fprintln(w, "Text 1")

		// Можно напрямую вызывать метод в интерфейсе ResponseWriter
		_, _ = w.Write([]byte("Text 2\n"))
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
