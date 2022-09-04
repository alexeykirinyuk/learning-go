package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)

		_, _ = fmt.Fprintln(w, "ok")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
