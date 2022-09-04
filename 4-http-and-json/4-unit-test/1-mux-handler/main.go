package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "hello")
}

func main() {
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", createRouter()))
}

func createRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	return router
}
