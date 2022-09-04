package main

import (
	"fmt"
	"log"
	"net/http"
)

type httpHandler struct {
	mux  *http.ServeMux
	name string
}

func newHttpHandler(name string) *httpHandler {
	handler := &httpHandler{
		mux:  &http.ServeMux{},
		name: name,
	}

	handler.mux.HandleFunc("/", handler.index)

	return handler
}

func (h httpHandler) index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "index "+h.name)
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func main() {
	h := newHttpHandler("alexey")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", h))
}
