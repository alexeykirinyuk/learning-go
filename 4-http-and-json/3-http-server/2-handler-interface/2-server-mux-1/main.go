package main

import "net/http"

func main() {
	// так не подходит, нужен указатель
	// var _ http.Handler = &http.ServeMux{}

	// можно так
	mux := http.NewServeMux()

	// или так
	mux = &http.ServeMux{} // все поля приватные

	// ну или так
	mux = new(http.ServeMux)

	_ = mux
}
