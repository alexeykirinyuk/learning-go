package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/heart", func(w http.ResponseWriter, r *http.Request) {
		// Без этой строки будет text/plain
		w.Header().Set("Content-Type", "application/json")

		_, _ = fmt.Fprint(w, `{"test":1}`)
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
