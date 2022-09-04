package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/heart", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1") // не может прочитать: тк content-length от ❤ = 3
		_, _ = fmt.Fprintln(w, "❤")
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
