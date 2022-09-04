package main

import (
	"fmt"
	"log"
	"net/http"
)

type HttpHandler struct {
}

// обрабатывает любой запрос GET/PUT/POST - по-любому пути
func (HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Был запрошен путь: %q", r.RequestURI)
}

func main() {
	const addr = "127.0.0.1:8080"
	httpHandler := HttpHandler{}
	server := http.Server{
		Addr:    addr,
		Handler: httpHandler,
	}

	log.Fatal(server.ListenAndServe())
	// то же самое что и
	// log.Fatal(http.ListenAndServer(addr, httpHandler))
}
